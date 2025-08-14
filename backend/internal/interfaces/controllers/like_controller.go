package controllers

import (
	"backend/pkg/utils"
	"net/http"

	"backend/internal/usecases"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LikeController interface {
	LikePost(c echo.Context) error
	UnlikePost(c echo.Context) error
	IsPostLikedByUser(c echo.Context) error
}

type likeController struct {
	likeService usecases.LikeUsecase
}

func NewLikeController(likeService usecases.LikeUsecase) LikeController {
	return &likeController{likeService: likeService}
}

func (h *likeController) LikePost(c echo.Context) error {
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

func (h *likeController) UnlikePost(c echo.Context) error {
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

func (h *likeController) IsPostLikedByUser(c echo.Context) error {
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
