package handler

import (
	"backend/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SpotifyHandler struct {
	spotifyService *service.SpotifyService
}

func NewSpotifyHandler(spotifyService *service.SpotifyService) *SpotifyHandler {
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
