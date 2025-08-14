package controllers

import (
	"backend/internal/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SpotifyController struct {
	spotifyService *usecases.SpotifyService
}

func NewSpotifyController(spotifyService *usecases.SpotifyService) *SpotifyController {
	return &SpotifyController{
		spotifyService: spotifyService,
	}
}

func (h *SpotifyController) SearchTracks(c echo.Context) error {
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
