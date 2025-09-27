package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid" // For generating unique message IDs
	"github.com/gorilla/websocket"
	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
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
	clients map[string]map[string]*Client

	broadcast chan entity.Message

	register chan *Client

	unregister chan *Client

	messageUsecase usecase.MessageUsecase
}

func NewHub(messageUsecase usecase.MessageUsecase) *Hub {
	return &Hub{
		broadcast:      make(chan entity.Message),
		register:       make(chan *Client),
		unregister:     make(chan *Client),
		clients:        make(map[string]map[string]*Client),
		messageUsecase: messageUsecase,
	}
}

// Hubのメインループ。接続の出入りとメッセージ配信を一元的に直列処理
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

type Client struct {
	id          string
	hub         *Hub
	conn        *websocket.Conn
	send        chan entity.Message
	communityID string
	userID      string
}

// WebSocketからの受信専用gorutineでクライアントの発言をHubに渡す
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
		msg := entity.Message{
			ID:          uuid.New().String(),
			CommunityID: c.communityID,
			SenderID:    c.userID,
			Content:     messageContent,
			CreatedAt:   time.Now(),
		}
		c.hub.broadcast <- msg
	}
}

// HubからClientへの送信用gorutine
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
			// Convert to API response shape to ensure proper JSON field names
			msgBytes, err := json.Marshal(response.ToMessageResponse(&message))
			if err != nil {
				log.Printf("Error marshalling message: %v", err)
				return
			}
			w.Write(msgBytes)

			// sendチャネルのキューをWebSocketに書き出し
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				m := <-c.send
				queuedMsgBytes, err := json.Marshal(response.ToMessageResponse(&m))
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

// HTTPリクエストをWebSocketにアップグレードし、Clientを作ってHubに登録、履歴を初期送信・read/write のゴルーチンを起動
func ServeWs(hub *Hub, messageUsecase usecase.MessageUsecase, w http.ResponseWriter, r *http.Request) {
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
