package handler

import (
	"net/http"

	 "github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type SpotifyHandler struct {
	spotifyUsecase *usecase.SpotifyUsecase
}

func NewSpotifyHandler(spotifyUsecase *usecase.SpotifyUsecase) *SpotifyHandler {
	return &SpotifyHandler{
		spotifyUsecase: spotifyUsecase,
	}
}

func (h *SpotifyHandler) SearchTracks(c echo.Context) error {
	query := c.QueryParam("q")

	if query == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing search term")
	}

	tracks, err := h.spotifyUsecase.SearchTracks(query)
	if err != nil {
		return err // spotifyUsecase.SearchTracks already returns echo.HTTPError
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"tracks":  tracks,
	})
}
