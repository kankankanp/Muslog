package handler

import (
	"net/http"

	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

// MessageHandler handles HTTP requests related to messages.
type MessageHandler struct {
	messageUsecase usecase.MessageUsecase
}

// NewMessageHandler creates a new MessageHandler.
func NewMessageHandler(messageUsecase usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{messageUsecase: messageUsecase}
}

// GetMessagesByCommunityID handles the retrieval of messages for a specific community.
func (mh *MessageHandler) GetMessagesByCommunityID(c echo.Context) error {
	communityID := c.Param("communityId")
	if communityID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Community ID is required"})
	}

	messages, err := mh.messageUsecase.GetMessagesByCommunityID(communityID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve messages"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"messages": messages})
}
