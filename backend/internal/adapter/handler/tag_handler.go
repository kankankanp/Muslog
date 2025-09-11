package handler

import (
	"net/http"
	"strconv"

	"github.com/kankankanp/Muslog/internal/adapter/dto/request"
	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TagHandler struct {
	tagUsecase usecase.TagUsecase
}

func NewTagHandler(tagUsecase usecase.TagUsecase) *TagHandler {
	return &TagHandler{tagUsecase: tagUsecase}
}

// =======================
// CRUD
// =======================

func (h *TagHandler) CreateTag(c echo.Context) error {
	var req request.TagRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	tag, err := h.tagUsecase.CreateTag(req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to create tag",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.TagDetailResponse{
		Message: "Tag created successfully",
		Tag:     response.ToTagResponse(tag),
	})
}

func (h *TagHandler) GetTagByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid tag ID",
		})
	}

	tag, err := h.tagUsecase.GetTagByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.CommonResponse{
			Message: "Tag not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.TagDetailResponse{
		Message: "Success",
		Tag:     response.ToTagResponse(tag),
	})
}

func (h *TagHandler) GetAllTags(c echo.Context) error {
	tags, err := h.tagUsecase.GetAllTags()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to fetch tags",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.TagListResponse{
		Message: "Success",
		Tags:    response.ToTagResponses(tags),
	})
}

func (h *TagHandler) UpdateTag(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid tag ID",
		})
	}

	var req request.TagRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	tag, err := h.tagUsecase.UpdateTag(uint(id), req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to update tag",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.TagDetailResponse{
		Message: "Tag updated successfully",
		Tag:     response.ToTagResponse(tag),
	})
}

func (h *TagHandler) DeleteTag(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid tag ID",
		})
	}

	if err := h.tagUsecase.DeleteTag(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to delete tag",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.CommonResponse{
		Message: "Tag deleted successfully",
	})
}

// =======================
// Post 関連
// =======================

func (h *TagHandler) AddTagsToPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("postID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid post ID",
		})
	}

	var req request.TagNamesRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	if err := h.tagUsecase.AddTagsToPost(uint(postID), req.TagNames); err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to add tags to post",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.CommonResponse{
		Message: "Tags added to post successfully",
	})
}

func (h *TagHandler) RemoveTagsFromPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("postID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid post ID",
		})
	}

	var req request.TagNamesRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	if err := h.tagUsecase.RemoveTagsFromPost(uint(postID), req.TagNames); err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to remove tags from post",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.CommonResponse{
		Message: "Tags removed from post successfully",
	})
}

func (h *TagHandler) GetTagsByPostID(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("postID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid post ID",
		})
	}

	tags, err := h.tagUsecase.GetTagsByPostID(uint(postID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to fetch tags by post",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.TagListResponse{
		Message: "Success",
		Tags:    response.ToTagResponses(tags),
	})
}
