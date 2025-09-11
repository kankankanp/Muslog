package handler

import (
	"net/http"

	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	Usecase usecase.MessageUsecase
}

func NewMessageHandler(u usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{Usecase: u}
}

func (h *MessageHandler) GetMessagesByCommunityID(c echo.Context) error {
	communityID := c.Param("communityId")
	if communityID == "" {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Community ID is required",
		})
	}

	messages, err := h.Usecase.GetMessagesByCommunityID(communityID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to retrieve messages",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.MessageListResponse{
		Message:  "Success",
		Messages: response.ToMessageResponses(messages),
	})
}
