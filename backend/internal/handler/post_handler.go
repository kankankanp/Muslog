package handler

import (
	"net/http"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	model "github.com/kankankanp/Muslog/internal/entity"
	service "github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	Service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{Service: service}
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

	posts, err := h.Service.GetAllPosts(userID)
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

	post, err := h.Service.GetPostByID(uint(id), userID)
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
	post := model.Post{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	for _, t := range req.Tracks {
		post.Tracks = append(post.Tracks, model.Track{
			SpotifyID:     t.SpotifyID,
			Name:          t.Name,
			ArtistName:    t.ArtistName,
			AlbumImageUrl: t.AlbumImageUrl,
		})
	}
	if err := h.Service.CreatePost(&post); err != nil {
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
	post, err := h.Service.GetPostByID(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	post.Title = req.Title
	post.Description = req.Description
	post.UpdatedAt = time.Now()
	if err := h.Service.UpdatePost(post); err != nil {
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
	if err := h.Service.DeletePost(uint(id)); err != nil {
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

	posts, totalCount, err := h.Service.GetPostsByPage(page, PerPage, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts, "totalCount": totalCount})
}
