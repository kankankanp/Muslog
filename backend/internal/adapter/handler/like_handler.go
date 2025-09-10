package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/kankankanp/Muslog/pkg/utils"
	"github.com/labstack/echo/v4"
)

type LikeHandler interface {
	LikePost(c echo.Context) error
	UnlikePost(c echo.Context) error
	IsPostLikedByUser(c echo.Context) error
}

type likeHandler struct {
	likeUsecase usecase.LikeUsecase
}

func (h *likeHandler) LikePost(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	liked, err := h.likeUsecase.ToggleLike(c.Request().Context(), postID, userID)
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

	if err := h.likeUsecase.UnlikePost(c.Request().Context(), postID, userID); err != nil {
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

	isLiked, err := h.likeUsecase.IsPostLikedByUser(c.Request().Context(), postID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"isLiked": isLiked})
}
