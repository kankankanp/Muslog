package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
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

// =======================
// Toggle Like
// =======================
func (h *LikeHandler) LikePost(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid post ID",
			Error:   err.Error(),
		})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	liked, err := h.Usecase.ToggleLike(c.Request().Context(), postID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to toggle like",
			Error:   err.Error(),
		})
	}

	if liked {
		return c.JSON(http.StatusOK, response.CommonResponse{
			Message: "Post liked successfully",
		})
	}
	return c.JSON(http.StatusOK, response.CommonResponse{
		Message: "Post unliked successfully",
	})
}

// =======================
// Unlike Post
// =======================
func (h *LikeHandler) UnlikePost(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid post ID",
			Error:   err.Error(),
		})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	if err := h.Usecase.UnlikePost(c.Request().Context(), postID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to unlike post",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.CommonResponse{
		Message: "Post unliked successfully",
	})
}

// =======================
// Check Like Status
// =======================
func (h *LikeHandler) IsPostLikedByUser(c echo.Context) error {
	postID, err := utils.ParseID(c, "postID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid post ID",
			Error:   err.Error(),
		})
	}

	userID := c.Get("userID").(string) // middleware でセットされる想定

	isLiked, err := h.Usecase.IsPostLikedByUser(c.Request().Context(), postID, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to check like status",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.LikeStatusResponse{
		IsLiked: isLiked,
	})
}

/*
ログインユーザーがいいねした投稿一覧を取得
*/
func (h *LikeHandler) GetLikedPostsByUser(c echo.Context) error {
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	posts, err := h.Usecase.GetLikedPostsByUser(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to get liked posts",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.PostListResponse{
		Message: "Success",
		Posts:   response.ToPostResponses(posts),
	})
}
