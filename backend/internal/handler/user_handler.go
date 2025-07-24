package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service *service.UserService
}

// Login godoc
// @Summary User login
// @Description Log in a user and return a JWT token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body model.User true "User credentials"
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func (h *UserHandler) Login(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request"})
	}

	user, err := h.Service.AuthenticateUser(u.Email, u.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthorized"})
	}

	// Create access token
	accessToken, err := createToken(user.ID, time.Hour*24) // 24 hours
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create access token"})
	}

	// Create refresh token
	refreshToken, err := createToken(user.ID, time.Hour*24*7) // 7 days
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create refresh token"})
	}

	// Set tokens in cookies
	setTokenCookie(c, "access_token", accessToken)
	setTokenCookie(c, "refresh_token", refreshToken)

	return c.JSON(http.StatusOK, echo.Map{"message": "Login successful", "user": user})
}

// RefreshToken godoc
// @Summary Refresh JWT token
// @Description Refresh the JWT access token using the refresh token.
// @Tags auth
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /refresh [post]
func (h *UserHandler) RefreshToken(c echo.Context) error {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Refresh token not found"})
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte("secret"), nil // Replace "secret" with your secret key
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid or expired refresh token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid token claims"})
	}

	userID := uint(claims["user_id"].(float64))

	// Create new access token
	accessToken, err := createToken(userID, time.Hour*24) // 24 hours
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Could not create access token"})
	}

	setTokenCookie(c, "access_token", accessToken)

	return c.JSON(http.StatusOK, echo.Map{"message": "Token refreshed"})
}

// GetMe godoc
// @Summary Get current user
// @Description Get the currently logged in user's information.
// @Tags auth
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /me [get]
func (h *UserHandler) GetMe(c echo.Context) error {
	userContext := c.Get("user").(jwt.MapClaims)
	userID := uint(userContext["user_id"].(float64))

	user, err := h.Service.GetUserByID(fmt.Sprintf("%d", userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "users": users})
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a single user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, err := h.Service.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not Found"})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "user": user})
}

// GetUserPosts godoc
// @Summary Get user posts
// @Description Get all posts by a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id}/posts [get]
func (h *UserHandler) GetUserPosts(c echo.Context) error {
	id := c.Param("id")
	posts, err := h.Service.GetUserPosts(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Error", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "posts": posts})
}

func createToken(userID uint, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret")) // Replace "secret" with a real secret key
}

func setTokenCookie(c echo.Context, name, token string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteLaxMode
	// cookie.Secure = true // In production, set this to true
	c.SetCookie(cookie)
}
 