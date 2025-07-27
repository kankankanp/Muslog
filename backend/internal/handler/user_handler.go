package handler

import (
	"fmt"
	"net/http"
	"simple-blog/backend/internal/model"
	"simple-blog/backend/internal/service"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) Login(c echo.Context) error {
	fmt.Println("Login handler called")
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	user, err := h.Service.AuthenticateUser(u.Email, u.Password)
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

	return c.JSON(http.StatusOK, echo.Map{"message": "Login successful", "user": user})
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

	user, err := h.Service.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Failed to register user"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully", "user": user})
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

	userID := uint(claims["user_id"].(float64))

	accessToken, err := createToken(fmt.Sprintf("%d", userID), time.Hour*24) // 24 hours
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
	userID := uint(userContext["user_id"].(float64))

	user, err := h.Service.GetUserByID(fmt.Sprintf("%d", userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "users": users})
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.Service.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "user": user})
}

func (h *UserHandler) GetUserPosts(c echo.Context) error {
	id := c.Param("id")
	posts, err := h.Service.GetUserPosts(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts})
}

func createToken(userID string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func setTokenCookie(c echo.Context, name, token string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	// cookie.Secure = true // 本番環境ではtrueにする
	c.SetCookie(cookie)
}

func clearTokenCookie(c echo.Context, name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	// cookie.Secure = true // 本番環境ではtrueにする
	c.SetCookie(cookie)
}
