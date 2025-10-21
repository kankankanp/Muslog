package handler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/adapter/dto/request"
	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type UserSettingHandler struct {
	Usecase usecase.UserSettingUsecase
}

func NewUserSettingHandler(u usecase.UserSettingUsecase) *UserSettingHandler {
	return &UserSettingHandler{Usecase: u}
}

func (h *UserSettingHandler) GetUserSetting(c echo.Context) error {
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	setting, err := h.Usecase.GetUserSetting(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to get user setting",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.UserSettingDetailResponse{
		Message: "Success",
		Setting: response.ToUserSettingResponse(setting),
	})
}

func (h *UserSettingHandler) UpdateUserSetting(c echo.Context) error {
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	var req request.UpdateUserSettingRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
	}

	setting, err := h.Usecase.UpdateUserSetting(c.Request().Context(), userID, req.EditorType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to update user setting",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.UserSettingDetailResponse{
		Message: "Setting updated successfully",
		Setting: response.ToUserSettingResponse(setting),
	})
}