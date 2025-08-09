package handler

import (
<<<<<<< HEAD
	"simple-blog/backend/internal/model"
	"simple-blog/backend/internal/service"
	"net/http"
	"strconv"
	"time"

=======
	"net/http"
	"simple-blog/backend/internal/model"
	"simple-blog/backend/internal/service"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
>>>>>>> develop
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
<<<<<<< HEAD
	Service service.PostService
}

func NewPostHandler(service service.PostService) *PostHandler {
=======
	Service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
>>>>>>> develop
	return &PostHandler{Service: service}
}

func (h *PostHandler) GetAllPosts(c echo.Context) error {
<<<<<<< HEAD
	posts, err := h.Service.GetAllPosts()
=======
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)
	posts, err := h.Service.GetAllPosts(userID)
>>>>>>> develop
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
<<<<<<< HEAD
	post, err := h.Service.GetPostByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "post": post, "likesCount": post.LikesCount})
}

func (h *PostHandler) CreatePost(c echo.Context) error {
=======
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)
	post, err := h.Service.GetPostByID(uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "post": post})
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

>>>>>>> develop
	type TrackInput struct {
		SpotifyID     string `json:"spotifyId"`
		Name          string `json:"name"`
		ArtistName    string `json:"artistName"`
		AlbumImageUrl string `json:"albumImageUrl"`
	}
	var req struct {
		Title       string       `json:"title"`
		Description string      `json:"description"`
<<<<<<< HEAD
		UserID      string      `json:"userId"`
=======
>>>>>>> develop
		Tracks      []TrackInput `json:"tracks"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request", "error": err.Error()})
	}
	post := model.Post{
		Title:       req.Title,
		Description: req.Description,
<<<<<<< HEAD
		UserID:      req.UserID,
=======
		UserID:      userID,
>>>>>>> develop
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
<<<<<<< HEAD
	post, err := h.Service.GetPostByID(uint(id))
=======
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)
	post, err := h.Service.GetPostByID(uint(id), userID)
>>>>>>> develop
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
<<<<<<< HEAD
	posts, totalCount, err := h.Service.GetPostsByPage(page, PerPage)
=======
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)
	posts, totalCount, err := h.Service.GetPostsByPage(page, PerPage, userID)
>>>>>>> develop
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts, "totalCount": totalCount})
}