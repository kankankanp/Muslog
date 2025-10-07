package handler

import (
	"errors"
	"log"
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

type BandRecruitmentHandler struct {
	Usecase usecase.BandRecruitmentUsecase
}

func NewBandRecruitmentHandler(usecase usecase.BandRecruitmentUsecase) *BandRecruitmentHandler {
	return &BandRecruitmentHandler{Usecase: usecase}
}

func (h *BandRecruitmentHandler) ListBandRecruitments(c echo.Context) error {
	filter := usecase.BandRecruitmentFilterInput{
		Keyword:  c.QueryParam("keyword"),
		Genre:    c.QueryParam("genre"),
		Location: c.QueryParam("location"),
		Status:   c.QueryParam("status"),
	}

	if page, err := strconv.Atoi(c.QueryParam("page")); err == nil {
		filter.Page = page
	}
	if perPage, err := strconv.Atoi(c.QueryParam("perPage")); err == nil {
		filter.PerPage = perPage
	}
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.PerPage == 0 {
		filter.PerPage = 10
	}

	userID := getUserIDFromContext(c)

	recruitments, total, err := h.Usecase.SearchBandRecruitments(c.Request().Context(), filter, userID)
	if err != nil {
		log.Printf("[BandRecruitments][List] error: %v (filter=%+v userID=%s)", err, filter, userID)
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{Message: "Error", Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response.BandRecruitmentListResponse{
		Message:      "Success",
		Recruitments: response.ToBandRecruitmentResponses(recruitments),
		TotalCount:   total,
		Page:         filter.Page,
		PerPage:      filter.PerPage,
	})
}

func (h *BandRecruitmentHandler) GetBandRecruitment(c echo.Context) error {
	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid ID"})
	}

	recruitment, err := h.Usecase.GetBandRecruitmentByID(c.Request().Context(), id, getUserIDFromContext(c))
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrRecruitmentNotFound):
			return c.JSON(http.StatusNotFound, response.CommonResponse{Message: "Not Found"})
		default:
			log.Printf("[BandRecruitments][Get] error: %v (id=%s)", err, id)
			return c.JSON(http.StatusInternalServerError, response.CommonResponse{Message: "Error", Error: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, response.BandRecruitmentDetailResponse{
		Message:     "Success",
		Recruitment: response.ToBandRecruitmentResponse(recruitment),
	})
}

func (h *BandRecruitmentHandler) CreateBandRecruitment(c echo.Context) error {
	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, response.CommonResponse{Message: "Unauthorized"})
	}
	userID, _ := claims["user_id"].(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, response.CommonResponse{Message: "Unauthorized"})
	}

	var req request.CreateBandRecruitmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid request", Error: err.Error()})
	}

	deadline, err := parseDeadline(req.Deadline)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid deadline", Error: err.Error()})
	}

	input := usecase.CreateBandRecruitmentInput{
		Title:           req.Title,
		Description:     req.Description,
		Genre:           req.Genre,
		Location:        req.Location,
		RecruitingParts: req.RecruitingParts,
		SkillLevel:      req.SkillLevel,
		Contact:         req.Contact,
		Deadline:        deadline,
		Status:          req.Status,
		UserID:          userID,
	}

	recruitment, err := h.Usecase.CreateBandRecruitment(c.Request().Context(), input)
	if err != nil {
		log.Printf("[BandRecruitments][Create] error: %v (userID=%s input=%+v)", err, userID, input)
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{Message: "Error", Error: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.BandRecruitmentDetailResponse{
		Message:     "Created",
		Recruitment: response.ToBandRecruitmentResponse(recruitment),
	})
}

func (h *BandRecruitmentHandler) UpdateBandRecruitment(c echo.Context) error {
	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, response.CommonResponse{Message: "Unauthorized"})
	}
	userID, _ := claims["user_id"].(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, response.CommonResponse{Message: "Unauthorized"})
	}

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid ID"})
	}

	var req request.UpdateBandRecruitmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid request", Error: err.Error()})
	}

	deadline, err := parseDeadline(req.Deadline)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid deadline", Error: err.Error()})
	}

	input := usecase.UpdateBandRecruitmentInput{
		ID:              id,
		Title:           req.Title,
		Description:     req.Description,
		Genre:           req.Genre,
		Location:        req.Location,
		RecruitingParts: req.RecruitingParts,
		SkillLevel:      req.SkillLevel,
		Contact:         req.Contact,
		Deadline:        deadline,
		Status:          req.Status,
		UserID:          userID,
	}

	recruitment, err := h.Usecase.UpdateBandRecruitment(c.Request().Context(), input)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrRecruitmentNoPrivilege):
			return c.JSON(http.StatusForbidden, response.CommonResponse{Message: "Forbidden"})
		case errors.Is(err, usecase.ErrRecruitmentNotFound):
			return c.JSON(http.StatusNotFound, response.CommonResponse{Message: "Not Found"})
		default:
			log.Printf("[BandRecruitments][Update] error: %v (input=%+v)", err, input)
			return c.JSON(http.StatusInternalServerError, response.CommonResponse{Message: "Error", Error: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, response.BandRecruitmentDetailResponse{
		Message:     "Updated",
		Recruitment: response.ToBandRecruitmentResponse(recruitment),
	})
}

func (h *BandRecruitmentHandler) ApplyToBandRecruitment(c echo.Context) error {
	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, response.CommonResponse{Message: "Unauthorized"})
	}
	userID, _ := claims["user_id"].(string)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, response.CommonResponse{Message: "Unauthorized"})
	}

	id := c.Param("id")
	if strings.TrimSpace(id) == "" {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid ID"})
	}

	var req request.ApplyBandRecruitmentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Invalid request", Error: err.Error()})
	}

	if err := h.Usecase.ApplyToBandRecruitment(c.Request().Context(), usecase.ApplyBandRecruitmentInput{
		BandRecruitmentID: id,
		ApplicantID:       userID,
		Message:           req.Message,
	}); err != nil {
		switch {
		case errors.Is(err, usecase.ErrRecruitmentNotFound):
			return c.JSON(http.StatusNotFound, response.CommonResponse{Message: "Not Found"})
		case errors.Is(err, usecase.ErrApplyOwnRecruitment):
			return c.JSON(http.StatusBadRequest, response.CommonResponse{Message: "Cannot apply to own recruitment"})
		case errors.Is(err, usecase.ErrRecruitmentClosed):
			return c.JSON(http.StatusConflict, response.CommonResponse{Message: "Recruitment closed"})
		case errors.Is(err, usecase.ErrAlreadyApplied):
			return c.JSON(http.StatusConflict, response.CommonResponse{Message: "Already applied"})
		default:
			log.Printf("[BandRecruitments][Apply] error: %v (id=%s applicant=%s)", err, id, userID)
			return c.JSON(http.StatusInternalServerError, response.CommonResponse{Message: "Error", Error: err.Error()})
		}
	}

	return c.JSON(http.StatusOK, response.CommonResponse{Message: "Applied"})
}

func parseDeadline(value *string) (*time.Time, error) {
	if value == nil {
		return nil, nil
	}
	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil, nil
	}

	if t, err := time.Parse(time.RFC3339, trimmed); err == nil {
		return &t, nil
	}
	if t, err := time.Parse("2006-01-02", trimmed); err == nil {
		return &t, nil
	}

	return nil, errors.New("invalid datetime format")
}

func getUserIDFromContext(c echo.Context) string {
	claims, ok := c.Get("user").(jwt.MapClaims)
	if !ok {
		return ""
	}
	userID, _ := claims["user_id"].(string)
	return userID
}
