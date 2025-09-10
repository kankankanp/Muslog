package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/kankankanp/Muslog/pkg/utils"
	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	Usecase usecase.LikeUsecase
}

func NewLikeHandler(u usecase.LikeUsecase) *LikeHandler {
	return &LikeHandler{Usecase: u}
}

func (h *LikeHandler) LikePost(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid post ID", "error": err.Error()})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	liked, err := h.Usecase.ToggleLike(c.Request().Context(), postID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to toggle like", "error": err.Error()})
	}

	if liked {
		return c.JSON(http.StatusOK, echo.Map{"message": "Post liked successfully"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Post unliked successfully"})
}

func (h *LikeHandler) UnlikePost(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid post ID", "error": err.Error()})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	if err := h.Usecase.UnlikePost(c.Request().Context(), postID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to unlike post", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Post unliked successfully"})
}

func (h *LikeHandler) IsPostLikedByUser(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid post ID", "error": err.Error()})
	}

	userID := c.Get("userID").(string) // middleware でセットされる想定

	isLiked, err := h.Usecase.IsPostLikedByUser(c.Request().Context(), postID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to check like status", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"isLiked": isLiked})
}
