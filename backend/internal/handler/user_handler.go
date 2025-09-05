package handler

import (
	"fmt"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v5"
	model "github.com/kankankanp/Muslog/internal/entity"
	service "github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *service.UserService
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func toUserResponse(user *model.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (h *UserHandler) Login(c echo.Context) error {
	fmt.Println("Login handler called")
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	user, err := h.Service.AuthenticateUser(c.Request().Context(), u.Email, u.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
	}

	accessToken, err := createToken(user.ID, time.Hour*24) // 有効期限は24時間
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create access token"})
	}

	refreshToken, err := createToken(user.ID, time.Hour*24*7) // 有効期限は7日間
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create refresh token"})
	}

	setTokenCookie(c, "access_token", accessToken)
	setTokenCookie(c, "refresh_token", refreshToken)

	return c.JSON(http.StatusOK, echo.Map{"message": "Login successful", "user": toUserResponse(user)})
}

func (h *UserHandler) Register(c echo.Context) error {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	user, err := h.Service.CreateUser(c.Request().Context(), req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to register user"})
	}

	accessToken, err := createToken(user.ID, time.Hour*24) // 有効期限は24時間
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create access token"})
	}

	refreshToken, err := createToken(user.ID, time.Hour*24*7) // 有効期限は7日間
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create refresh token"})
	}

	setTokenCookie(c, "access_token", accessToken)
	setTokenCookie(c, "refresh_token", refreshToken)

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully", "user": toUserResponse(user)})
}

func (h *UserHandler) RefreshToken(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Refresh token not found"})
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid or expired refresh token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid token claims"})
	}

	userID := claims["user_id"].(string)

	accessToken, err := createToken(userID, time.Hour*24) // 24 hours
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create access token"})
	}

	setTokenCookie(c, "access_token", accessToken)

	return c.JSON(http.StatusOK, echo.Map{"message": "Token refreshed", "accessToken": accessToken})
}

func (h *UserHandler) Logout(c echo.Context) error {
	clearTokenCookie(c, "access_token")
	clearTokenCookie(c, "refresh_token")
	return c.JSON(http.StatusOK, echo.Map{"message": "Logout successful"})
}

func (h *UserHandler) GetMe(c echo.Context) error {
	userContext := c.Get("user").(jwt.MapClaims)
	userID := userContext["user_id"].(string)

	user, err := h.Service.GetUserByID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}

	return c.JSON(http.StatusOK, toUserResponse(user))
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}

	// パスワードを除外したユーザー一覧を作成
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, toUserResponse(&user))
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "users": userResponses})
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.Service.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "user": toUserResponse(user)})
}

func (h *UserHandler) GetUserPosts(c echo.Context) error {
	id := c.Param("id")
	posts, err := h.Service.GetUserPosts(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts})
}