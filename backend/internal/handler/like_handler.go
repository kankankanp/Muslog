package handler

import (
	"net/http"
	"simple-blog/backend/internal/service"
	"simple-blog/backend/pkg/utils"

	"github.com/labstack/echo/v4"
)

type LikeHandler interface {
	LikePost(c echo.Context) error
	UnlikePost(c echo.Context) error
	IsPostLikedByUser(c echo.Context) error
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

	userID := c.Get("userID").(string) // From auth middleware

	if err := h.likeService.LikePost(postID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Post liked successfully"})
}

func (h *likeHandler) UnlikePost(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	userID := c.Get("userID").(string) // From auth middleware

	if err := h.likeService.UnlikePost(postID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Post unliked successfully"})
}

func (h *likeHandler) IsPostLikedByUser(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	userID := c.Get("userID").(string) // From auth middleware

	isLiked, err := h.likeService.IsPostLikedByUser(postID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"isLiked": isLiked})
}
