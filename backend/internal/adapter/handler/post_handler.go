package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	Usecase usecase.PostUsecase
}

func NewPostHandler(usecase usecase.PostUsecase) *PostHandler {
	return &PostHandler{Usecase: usecase}
}

func (h *PostHandler) GetAllPosts(c echo.Context) error {
	var userID string
	userContext := c.Get("user")
	if userContext != nil {
		claims, ok := userContext.(jwt.MapClaims)
		if ok {
			userID, _ = claims["user_id"].(string)
		}
	}

	posts, err := h.Usecase.GetAllPosts(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts})
}

func (h *PostHandler) GetPostByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ID"})
	}

	var userID string
	userContext := c.Get("user")
	if userContext != nil {
		claims, ok := userContext.(jwt.MapClaims)
		if ok {
			userID, _ = claims["user_id"].(string)
		}
	}

	post, err := h.Usecase.GetPostByID(c.Request().Context(), uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "post": post})
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	type TrackInput struct {
		SpotifyID     string `json:"spotifyId"`
		Name          string `json:"name"`
		ArtistName    string `json:"artistName"`
		AlbumImageUrl string `json:"albumImageUrl"`
	}
	var req struct {
		Title       string       `json:"title"`
		Description string       `json:"description"`
		Tracks      []TrackInput `json:"tracks"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request", "error": err.Error()})
	}
	post := entity.Post{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	for _, t := range req.Tracks {
		post.Tracks = append(post.Tracks, entity.Track{
			SpotifyID:     t.SpotifyID,
			Name:          t.Name,
			ArtistName:    t.ArtistName,
			AlbumImageUrl: t.AlbumImageUrl,
		})
	}
	if err := h.Usecase.CreatePost(c.Request().Context(), &post); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Success", "post": post})
}

func (h *PostHandler) UpdatePost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ID"})
	}
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request", "error": err.Error()})
	}
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)
	post, err := h.Usecase.GetPostByID(c.Request().Context(), uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	post.Title = req.Title
	post.Description = req.Description
	post.UpdatedAt = time.Now()
	if err := h.Usecase.UpdatePost(c.Request().Context(), post); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "post": post})
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ID"})
	}
	if err := h.Usecase.DeletePost(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success"})
}

func (h *PostHandler) GetPostsByPage(c echo.Context) error {
	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid page"})
	}
	const PerPage = 4

	var userID string
	userContext := c.Get("user")
	if userContext != nil {
		claims, ok := userContext.(jwt.MapClaims)
		if ok {
			userID, _ = claims["user_id"].(string)
		}
	}

	posts, totalCount, err := h.Usecase.GetPostsByPage(c.Request().Context(), page, PerPage, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts, "totalCount": totalCount})
}

// SearchPosts handles searching for posts.
func (h *PostHandler) SearchPosts(c echo.Context) error {
	query := c.QueryParam("q")
	tagsStr := c.QueryParam("tags")
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("perPage")

	var tags []string
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1 // Default to page 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10 // Default to 10 items per page
	}

	var userID string
	userContext := c.Get("user")
	if userContext != nil {
		claims, ok := userContext.(jwt.MapClaims)
		if ok {
			userID, _ = claims["user_id"].(string)
		}
	}

	posts, totalCount, err := h.Usecase.SearchPosts(c.Request().Context(), query, tags, page, perPage, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to search posts", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Posts search successful",
		"posts":      posts,
		"totalCount": totalCount,
		"page":       page,
		"perPage":    perPage,
	})
}
