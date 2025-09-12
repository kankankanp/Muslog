package handler

import (
	"net/http"

	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/usecase"
	"github.com/labstack/echo/v4"
)

type SpotifyHandler struct {
	spotifyUsecase usecase.SpotifyUsecase
}

func NewSpotifyHandler(spotifyUsecase usecase.SpotifyUsecase) *SpotifyHandler {
	return &SpotifyHandler{
		spotifyUsecase: spotifyUsecase,
	}
}

func (h *SpotifyHandler) SearchTracks(c echo.Context) error {
	query := c.QueryParam("q")

	if query == "" {
		return c.JSON(http.StatusBadRequest, response.CommonResponse{
			Message: "Missing search term",
		})
	}

	tracks, err := h.spotifyUsecase.SearchTracks(query)
	if err != nil {
		// Usecase 側で echo.HTTPError を返しているならそのまま
		return err
	}

	return c.JSON(http.StatusOK, response.SearchTracksResponse{
		Message: "Success",
		Tracks:  response.ToTrackResponses(tracks),
	})
}
