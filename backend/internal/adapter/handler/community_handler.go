package handler

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/adapter/dto/request"
	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type CommunityHandler struct {
	communityUsecase usecase.CommunityUsecase
}

func NewCommunityHandler(communityUsecase usecase.CommunityUsecase) *CommunityHandler {
	return &CommunityHandler{communityUsecase: communityUsecase}
}

// =======================
// Create Community
// =======================
func (h *CommunityHandler) CreateCommunity(c echo.Context) error {
	var req request.CreateCommunityRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	userContext := c.Get("user").(jwt.MapClaims)
	creatorID := userContext["user_id"].(string)

	community, err := h.communityUsecase.CreateCommunity(c.Request().Context(), req.Name, req.Description, creatorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to create community",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.CreateCommunityResponse{
		Message:   "Community created successfully",
		Community: response.ToCommunityResponse(community),
	})
}

// =======================
// Get All Communities
// =======================
func (h *CommunityHandler) GetAllCommunities(c echo.Context) error {
	communities, err := h.communityUsecase.GetAllCommunities(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to retrieve communities",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.CommunityListResponse{
		Message:     "Communities retrieved successfully",
		Communities: response.ToCommunityResponses(communities),
	})
}

// =======================
// Search Communities
// =======================
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
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to search communities",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.CommunityListResponse{
		Message:     "Communities search successful",
		Communities: response.ToCommunityResponses(communities),
		TotalCount:  totalCount,
		Page:        page,
		PerPage:     perPage,
	})
}
