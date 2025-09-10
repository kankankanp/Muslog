package handler

import (
	"net/http"

	service "github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type SpotifyHandler struct {
	spotifyService *service.SpotifyUsecase
}

func NewSpotifyHandler(spotifyService *service.SpotifyUsecase) *SpotifyHandler {
	return &SpotifyHandler{
		spotifyService: spotifyService,
	}
}

func (h *SpotifyHandler) SearchTracks(c echo.Context) error {
	query := c.QueryParam("q")

	if query == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing search term")
	}

	tracks, err := h.spotifyService.SearchTracks(query)
	if err != nil {
		return err // spotifyService.SearchTracks already returns echo.HTTPError
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"tracks":  tracks,
	})
}
