package handler

import (
	"net/http"
	"strconv"

	service "github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TagHandler struct {
	tagService service.TagUsecase
}

func NewTagHandler(tagService service.TagUsecase) *TagHandler {
	return &TagHandler{tagService: tagService}
}

func (h *TagHandler) CreateTag(c echo.Context) error {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tag, err := h.tagService.CreateTag(req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, tag)
}

func (h *TagHandler) GetTagByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tag ID")
	}

	tag, err := h.tagService.GetTagByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Tag not found")
	}

	return c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) GetAllTags(c echo.Context) error {
	tags, err := h.tagService.GetAllTags()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tags)
}

func (h *TagHandler) UpdateTag(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tag ID")
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tag, err := h.tagService.UpdateTag(uint(id), req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tag)
}

func (h *TagHandler) DeleteTag(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid tag ID")
	}

	if err := h.tagService.DeleteTag(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TagHandler) AddTagsToPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("postID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	var req struct {
		TagNames []string `json:"tag_names"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.tagService.AddTagsToPost(uint(postID), req.TagNames); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TagHandler) RemoveTagsFromPost(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("postID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	var req struct {
		TagNames []string `json:"tag_names"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.tagService.RemoveTagsFromPost(uint(postID), req.TagNames); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TagHandler) GetTagsByPostID(c echo.Context) error {
	postID, err := strconv.ParseUint(c.Param("postID"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid post ID")
	}

	tags, err := h.tagService.GetTagsByPostID(uint(postID))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tags)
}
