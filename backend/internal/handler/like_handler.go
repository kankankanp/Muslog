package handler

import (
	"net/http"
	"simple-blog/backend/internal/service"
	"simple-blog/backend/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LikeHandler interface {
	LikePost(c echo.Context) error
	UnlikePost(c echo.Context) error
}

type likeHandler struct {
	likeService service.LikeService
}

func NewLikeHandler(likeService service.LikeService) LikeHandler {
	return &likeHandler{likeService: likeService}
}

func (h *likeHandler) LikePost(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	liked, err := h.likeService.ToggleLike(postID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if liked {
		return c.JSON(http.StatusOK, echo.Map{"message": "Post liked successfully"})
	} else {
		return c.JSON(http.StatusOK, echo.Map{"message": "Post unliked successfully"})
	}
}

func (h *likeHandler) UnlikePost(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	if err := h.likeService.UnlikePost(postID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Post unliked successfully"})
}
