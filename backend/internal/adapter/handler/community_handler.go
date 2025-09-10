package handler

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type CommunityHandler struct {
	communityUsecase usecase.CommunityUsecase
}

func NewCommunityHandler(communityUsecase usecase.CommunityUsecase) *CommunityHandler {
	return &CommunityHandler{communityUsecase: communityUsecase}
}

func (h *CommunityHandler) CreateCommunity(c echo.Context) error {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	creatorID := userContext["user_id"].(string)

	community, err := h.communityUsecase.CreateCommunity(c.Request().Context(), req.Name, req.Description, creatorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create community"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Community created successfully", "community": community})
}

func (h *CommunityHandler) GetAllCommunities(c echo.Context) error {
	communities, err := h.communityUsecase.GetAllCommunities(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve communities"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Communities retrieved successfully", "communities": communities})
}

func (h *CommunityHandler) SearchCommunities(c echo.Context) error {
	query := c.QueryParam("q")
	pageStr := c.QueryParam("page")
	perPageStr := c.QueryParam("perPage")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10
	}

	communities, totalCount, err := h.communityUsecase.SearchCommunities(c.Request().Context(), query, page, perPage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to search communities", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Communities search successful",
		"communities": communities,
		"totalCount":  totalCount,
		"page":        page,
		"perPage":     perPage,
	})
}
