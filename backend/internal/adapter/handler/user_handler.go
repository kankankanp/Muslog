package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/kankankanp/Muslog/pkg/utils"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{Usecase: u}
}

func (h *UserHandler) GetMe(c echo.Context) error {
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	user, err := h.Usecase.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}

	return c.JSON(http.StatusOK, response.ToUserResponse(user))
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.Usecase.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.ToUserResponse(&user))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"users":   userResponses,
	})
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.Usecase.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"user":    response.ToUserResponse(user),
	})
}

func (h *UserHandler) GetUserPosts(c echo.Context) error {
	id := c.Param("id")
	posts, err := h.Usecase.GetUserPosts(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success",
		"posts":   posts,
	})
}

func (h *UserHandler) Login(c echo.Context) error {
	u := new(entity.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	user, err := h.Usecase.AuthenticateUser(c.Request().Context(), u.Email, u.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unauthorized",
			"error":   err.Error(),
		})
	}

	accessToken, err := utils.CreateToken(user.ID, time.Hour*24)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create access token"})
	}

	refreshToken, err := utils.CreateToken(user.ID, time.Hour*24*7)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create refresh token"})
	}

	utils.SetTokenCookie(c, "access_token", accessToken)
	utils.SetTokenCookie(c, "refresh_token", refreshToken)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"user":    response.ToUserResponse(user),
	})
}

func (h *UserHandler) Register(c echo.Context) error {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	user, err := h.Usecase.CreateUser(c.Request().Context(), req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to register user",
			"error":   err.Error(),
		})
	}

	accessToken, err := utils.CreateToken(user.ID, time.Hour*24)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create access token"})
	}

	refreshToken, err := utils.CreateToken(user.ID, time.Hour*24*7)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create refresh token"})
	}

	utils.SetTokenCookie(c, "access_token", accessToken)
	utils.SetTokenCookie(c, "refresh_token", refreshToken)

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"user":    response.ToUserResponse(user),
	})
}

func (h *UserHandler) RefreshToken(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Refresh token not found",
			"error":   err.Error(),
		})
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte("secret"), nil // TODO: 環境変数に置き換える
	})
	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid or expired refresh token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid token claims"})
	}

	userID := claims["user_id"].(string)

	accessToken, err := utils.CreateToken(userID, time.Hour*24)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create access token"})
	}

	utils.SetTokenCookie(c, "access_token", accessToken)

	return c.JSON(http.StatusOK, echo.Map{
		"message":     "Token refreshed",
		"accessToken": accessToken,
	})
}

func (h *UserHandler) Logout(c echo.Context) error {
	utils.ClearTokenCookie(c, "access_token")
	utils.ClearTokenCookie(c, "refresh_token")
	return c.JSON(http.StatusOK, echo.Map{"message": "Logout successful"})
}
