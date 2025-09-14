package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/adapter/dto/request"
	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	Usecase    usecase.PostUsecase
	TagUsecase usecase.TagUsecase
}

func NewPostHandler(usecase usecase.PostUsecase, tagUsecase usecase.TagUsecase) *PostHandler {
	return &PostHandler{Usecase: usecase, TagUsecase: tagUsecase}
}

func (h *PostHandler) GetAllPosts(c echo.Context) error {
	var userID string
	if claims, ok := c.Get("user").(jwt.MapClaims); ok {
		userID, _ = claims["user_id"].(string)
	}

	posts, err := h.Usecase.GetAllPosts(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Error", Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.PostListResponse{
		Message: "Success",
		Posts:   response.ToPostResponses(posts),
	})
}

func (h *PostHandler) GetPostByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid ID"})
	}

	var userID string
	if claims, ok := c.Get("user").(jwt.MapClaims); ok {
		userID, _ = claims["user_id"].(string)
	}

	post, err := h.Usecase.GetPostByID(c.Request().Context(), uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.CommonResponse{Message: "Not Found"})
	}

	return c.JSON(http.StatusOK, response.PostDetailResponse{
		Message: "Success",
		Post:    response.ToPostResponse(post),
	})
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	claims := c.Get("user").(jwt.MapClaims)
	userID := claims["user_id"].(string)

	var req request.CreatePostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid request", Error: err.Error(),
		})
	}

	input := usecase.CreatePostInput{
		Title:          req.Title,
		Description:    req.Description,
		UserID:         userID,
		HeaderImageUrl: req.HeaderImageUrl,
		Tags:           req.Tags,
	}

	for _, t := range req.Tracks {
		input.Tracks = append(input.Tracks, usecase.TrackInput{
			SpotifyID:     t.SpotifyID,
			Name:          t.Name,
			ArtistName:    t.ArtistName,
			AlbumImageUrl: t.AlbumImageUrl,
		})
	}

	post, err := h.Usecase.CreatePost(c.Request().Context(), input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Error", Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ToPostResponse(post))
}

func (h *PostHandler) UpdatePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid ID"})
	}

	var req request.UpdatePostRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid request", Error: err.Error(),
		})
	}

	claims := c.Get("user").(jwt.MapClaims)
	userID := claims["user_id"].(string)

	post, err := h.Usecase.GetPostByID(c.Request().Context(), uint(id), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.CommonResponse{Message: "Not Found"})
	}

	post.Title = req.Title
	post.Description = req.Description
	post.UpdatedAt = time.Now()

	if err := h.Usecase.UpdatePost(c.Request().Context(), post); err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Error", Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.PostDetailResponse{
		Message: "Success",
		Post:    response.ToPostResponse(post),
	})
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid ID"})
	}

	if err := h.Usecase.DeletePost(c.Request().Context(), uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Error", Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.CommonResponse{Message: "Success"})
}

func (h *PostHandler) GetPostsByPage(c echo.Context) error {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil || page < 1 {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid page"})
	}
	const PerPage = 4

	var userID string
	if claims, ok := c.Get("user").(jwt.MapClaims); ok {
		userID, _ = claims["user_id"].(string)
	}

	posts, totalCount, err := h.Usecase.GetPostsByPage(c.Request().Context(), page, PerPage, userID)
	if err != nil {
		// クライアント切断やタイムアウトの場合は 500 を返さず静かに終了
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) || errors.Is(c.Request().Context().Err(), context.Canceled) {
			return nil
		}
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Error", Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.PostListResponse{
		Message:    "Success",
		Posts:      response.ToPostResponses(posts),
		TotalCount: totalCount,
		Page:       page,
		PerPage:    PerPage,
	})
}

func (h *PostHandler) SearchPosts(c echo.Context) error {
	query := c.QueryParam("q")
	tagsStr := c.QueryParam("tags")
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("perPage")

	var tags []string
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}

	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	perPage, _ := strconv.Atoi(perPageStr)
	if perPage < 1 {
		perPage = 10
	}

	var userID string
	if claims, ok := c.Get("user").(jwt.MapClaims); ok {
		userID, _ = claims["user_id"].(string)
	}

	posts, totalCount, err := h.Usecase.SearchPosts(c.Request().Context(), query, tags, page, perPage, userID)
	if err != nil {
		// クライアント切断やタイムアウトの場合は 500 を返さず静かに終了
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) || errors.Is(c.Request().Context().Err(), context.Canceled) {
			return nil
		}
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to search posts", Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.PostListResponse{
		Message:    "Posts search successful",
		Posts:      response.ToPostResponses(posts),
		TotalCount: totalCount,
		Page:       page,
		PerPage:    perPage,
	})
}
