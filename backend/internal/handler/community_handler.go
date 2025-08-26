package handler

import (
	"net/http"

	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

// CommunityHandler handles HTTP requests for communities.
type CommunityHandler struct {
	communityUsecase usecase.CommunityUsecase
}

// NewCommunityHandler creates a new CommunityHandler.
func NewCommunityHandler(communityUsecase usecase.CommunityUsecase) *CommunityHandler {
	return &CommunityHandler{communityUsecase: communityUsecase}
}

// CreateCommunity handles the creation of a new community.
func (h *CommunityHandler) CreateCommunity(c echo.Context) error {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} // creatorID will be extracted from auth context

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// TODO: Extract actual creatorID from authenticated user context
	// For now, using a placeholder or a default user ID.
	creatorID := "placeholder_creator_id" // Replace with actual user ID from JWT/session

	community, err := h.communityUsecase.CreateCommunity(req.Name, req.Description, creatorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create community"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Community created successfully", "community": community})
}

// GetAllCommunities handles retrieving all communities.
func (h *CommunityHandler) GetAllCommunities(c echo.Context) error {
	communities, err := h.communityUsecase.GetAllCommunities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve communities"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Communities retrieved successfully", "communities": communities})
}
