package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/kankankanp/Muslog/pkg/utils"
	"github.com/labstack/echo/v4"
)

type OAuthHandler struct {
	Usecase usecase.OAuthUsecase
}

func NewOAuthHandler(usecase usecase.OAuthUsecase) *OAuthHandler {
	return &OAuthHandler{Usecase: usecase}
}

// Google 認証 URL を返す
func (h *OAuthHandler) GetGoogleAuthURL(c echo.Context) error {
	state := generateRandomState()
	setStateCookie(c, state)

	authURL := h.Usecase.GetAuthURL(state)

	return c.JSON(http.StatusOK, response.GoogleAuthURLResponse{
		AuthURL: authURL,
	})
}

// Google コールバック処理
func (h *OAuthHandler) GoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	stateCookie, err := c.Cookie("oauth_state")
	if err != nil || stateCookie.Value != state {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Invalid state parameter",
		})
	}

	clearStateCookie(c)

	if code == "" {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Authorization code not provided",
		})
	}

	user, err := h.Usecase.HandleCallback(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Failed to authenticate with Google",
			Error:   err.Error(),
		})
	}

	accessToken, err := utils.CreateToken(user.ID, time.Hour*24)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Could not create access token",
		})
	}

	refreshToken, err := utils.CreateToken(user.ID, time.Hour*24*7)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.CommonResponse{
			Message: "Could not create refresh token",
		})
	}

	utils.SetTokenCookie(c, "access_token", accessToken)
	utils.SetTokenCookie(c, "refresh_token", refreshToken)

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	// 認証後はフロントにリダイレクト
	return c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/dashboard")
}

// =======================
// Cookie & State Utils
// =======================

func generateRandomState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func setStateCookie(c echo.Context, state string) {
	cookie := &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		Expires:  time.Now().Add(time.Minute * 10),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(cookie)
}

func clearStateCookie(c echo.Context) {
	cookie := &http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	c.SetCookie(cookie)
}
