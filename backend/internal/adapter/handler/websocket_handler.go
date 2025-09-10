package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid" // For generating unique message IDs
	"github.com/gorilla/websocket"
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/kankankanp/Muslog/internal/usecase"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 300 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var (
	newline = []byte{' '}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for simplicity during development.
		// In production, you should restrict this to your frontend's origin.
		return true
	},
}

type Hub struct {
	// Registered clients for each community.
	// communityID -> map[clientID]*Client
	clients map[string]map[string]*Client

	// Inbound messages from the clients.
	broadcast chan entity.Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Message usecase for saving and retrieving messages.
	messageUsecase usecase.MessageUsecase // New field
}

// NewHub creates and returns a new Hub instance.
func NewHub(messageUsecase usecase.MessageUsecase) *Hub {
	return &Hub{
		broadcast:      make(chan entity.Message),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		clients:        make(map[string]map[string]*Client),
		messageUsecase: messageUsecase, // Assign new field
	}
}

// Run starts the hub, listening for events on its channels.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			if _, ok := h.clients[client.communityID]; !ok {
				h.clients[client.communityID] = make(map[string]*Client)
			}
			h.clients[client.communityID][client.id] = client
			log.Printf("Client %s registered to community %s. Total clients in community: %d", client.id, client.communityID, len(h.clients[client.communityID]))
		case client := <-h.unregister:
			if _, ok := h.clients[client.communityID]; ok {
				if _, ok := h.clients[client.communityID][client.id]; ok {
					delete(h.clients[client.communityID], client.id)
					close(client.send)
					log.Printf("Client %s unregistered from community %s. Total clients in community: %d", client.id, client.communityID, len(h.clients[client.communityID]))
					if len(h.clients[client.communityID]) == 0 {
						delete(h.clients, client.communityID)
						log.Printf("Community %s has no active clients, removing from hub.", client.communityID)
					}
				}
			}
		case message := <-h.broadcast:
			// Save message to database
			if err := h.messageUsecase.SaveMessage(&message); err != nil {
				log.Printf("Error saving message to DB: %v", err)
				// Optionally, handle error by not broadcasting or notifying sender
			}

			log.Printf("Broadcasting message to community %s: %s", message.CommunityID, message.Content)
			if clients, ok := h.clients[message.CommunityID]; ok {
				for _, client := range clients {
					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients[client.communityID], client.id)
					}
				}
			}
		}
	}
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	id          string
	hub         *Hub
	conn        *websocket.Conn
	send        chan entity.Message // Buffered channel of outbound messages.
	communityID string
	userID      string // Assuming a user ID for the sender
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a goroutine for each connection. The
// application ensures that there is at most one reader on a connection by
// executing all reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, messageBytes, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		messageContent := string(messageBytes)
		// For simplicity, assuming the message from client is just the content string.
		// In a real app, you might expect a JSON object with sender info etc.
		msg := entity.Message{
			ID:          uuid.New().String(),
			CommunityID: c.communityID,
			SenderID:    c.userID, // Use the client's user ID
			Content:     messageContent,
			CreatedAt:   time.Now(),
		}
		c.hub.broadcast <- msg
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			msgBytes, err := json.Marshal(message)
			if err != nil {
				log.Printf("Error marshalling message: %v", err)
				return
			}
			w.Write(msgBytes)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				queuedMsgBytes, err := json.Marshal(<-c.send)
				if err != nil {
					log.Printf("Error marshalling queued message: %v", err)
					continue
				}
				w.Write(queuedMsgBytes)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func ServeWs(hub *Hub, messageUsecase usecase.MessageUsecase, w http.ResponseWriter, r *http.Request) {
	// Extract community ID from URL path.
	// Assuming URL format like /ws/community/{communityId}
	communityID := r.URL.Path[len("/ws/community/"):]
	if communityID == "" {
		http.Error(w, "Community ID is required", http.StatusBadRequest)
		return
	}

	// TODO: In a real application, you would extract the userID from the request context
	// (e.g., from a JWT token after authentication).
	// For now, we'll use a placeholder or generate a random one.
	// For demonstration, let's use a simple placeholder.
	userID := "guest_user_" + uuid.New().String()[:8] // Placeholder user ID

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		id:          uuid.New().String(), // Unique ID for this client connection
		hub:         hub,
		conn:        conn,
		send:        make(chan entity.Message, 256),
		communityID: communityID,
		userID:      userID,
	}
	client.hub.register <- client

	// Send historical messages to the newly connected client
	messages, err := messageUsecase.GetMessagesByCommunityID(communityID)
	if err != nil {
		log.Printf("Error getting historical messages for community %s: %v", communityID, err)
	} else {
		for _, msg := range messages {
			select {
			case client.send <- msg:
			default:
				// If client's send buffer is full, skip sending historical messages
				log.Printf("Client %s send buffer full, skipping historical message: %+v", client.id, msg)
			}
		}
	}

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
