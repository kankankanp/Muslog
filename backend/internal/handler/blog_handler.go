package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type BlogHandler struct {
	Service *service.BlogService
}

// GetAllBlogs godoc
// @Summary Get all blogs
// @Description Get all blog posts
// @Tags blogs
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /blogs [get]
func (h *BlogHandler) GetAllBlogs(c echo.Context) error {
	posts, err := h.Service.GetAllBlogs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts})
}

// GetBlogByID godoc
// @Summary Get a blog by ID
// @Description Get a single blog post by its ID
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Success 200 {object} map[string]interface{}
// @Router /blogs/{id} [get]
func (h *BlogHandler) GetBlogByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ID"})
	}
	post, err := h.Service.GetBlogByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "post": post})
}

// CreateBlog godoc
// @Summary Create a new blog
// @Description Create a new blog post
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param blog body model.Post true "Blog post to create"
// @Success 201 {object} map[string]interface{}
// @Router /blogs [post]
func (h *BlogHandler) CreateBlog(c echo.Context) error {
	type TrackInput struct {
		SpotifyID     string `json:"spotifyId"`
		Name          string `json:"name"`
		ArtistName    string `json:"artistName"`
		AlbumImageUrl string `json:"albumImageUrl"`
	}
	var req struct {
		Title       string       `json:"title"`
		Description string      `json:"description"`
		UserID      string      `json:"userId"`
		Tracks      []TrackInput `json:"tracks"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request", "error": err.Error()})
	}
	post := model.Post{
		Title:       req.Title,
		Description: req.Description,
		UserID:      req.UserID,
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
	if err := h.Service.CreateBlog(&post); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Success", "post": post})
}

// UpdateBlog godoc
// @Summary Update a blog
// @Description Update an existing blog post
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Param blog body model.Post true "Blog post to update"
// @Success 200 {object} map[string]interface{}
// @Router /blogs/{id} [put]
func (h *BlogHandler) UpdateBlog(c echo.Context) error {
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
	post, err := h.Service.GetBlogByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	post.Title = req.Title
	post.Description = req.Description
	post.UpdatedAt = time.Now()
	if err := h.Service.UpdateBlog(post); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "post": post})
}

// DeleteBlog godoc
// @Summary Delete a blog
// @Description Delete a blog post by its ID
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param id path int true "Blog ID"
// @Success 200 {object} map[string]interface{}
// @Router /blogs/{id} [delete]
func (h *BlogHandler) DeleteBlog(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid ID"})
	}
	if err := h.Service.DeleteBlog(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success"})
}

// GetBlogsByPage godoc
// @Summary Get blogs by page
// @Description Get blog posts paginated
// @Tags blogs
// @Accept  json
// @Produce  json
// @Param page path int true "Page number"
// @Success 200 {object} map[string]interface{}
// @Router /blogs/page/{page} [get]
func (h *BlogHandler) GetBlogsByPage(c echo.Context) error {
	pageStr := c.Param("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid page"})
	}
	const PerPage = 4
	posts, totalCount, err := h.Service.GetBlogsByPage(page, PerPage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts, "totalCount": totalCount})
} 
