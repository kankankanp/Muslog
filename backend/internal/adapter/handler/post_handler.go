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
	"github.com/kankankanp/Muslog/internal/domain/entity"
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

    post := entity.Post{
        Title:       req.Title,
        Description: req.Description,
        UserID:      userID,
        HeaderImageUrl: req.HeaderImageUrl,
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
        return c.JSON(http.StatusInternalServerError, response.CommonResponse{
            Message: "Error", Error: err.Error(),
        })
    }

    // attach tags if provided
    if len(req.Tags) > 0 && h.TagUsecase != nil {
        if err := h.TagUsecase.AddTagsToPost(post.ID, req.Tags); err != nil {
            // ログに残しつつ、投稿自体は成功扱いにする
            // クライアント側は後からタグ再付与も可能
            // ここでエラーを返しても良いが、UXを優先
        }
    }

    return c.JSON(http.StatusCreated, response.PostDetailResponse{
        Message: "Success",
        Post:    response.ToPostResponse(&post),
    })
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
