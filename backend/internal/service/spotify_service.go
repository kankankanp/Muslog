package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type SpotifyTrackResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Album struct {
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	} `json:"album"`
}

type SpotifySearchResponse struct {
	Tracks struct {
		Items []SpotifyTrackResponse `json:"items"`
	} `json:"tracks"`
}

type FormattedTrack struct {
	SpotifyID     string `json:"spotifyId"`
	Name          string `json:"name"`
	ArtistName    string `json:"artistName"`
	AlbumImageURL string `json:"albumImageUrl"`
}

type SpotifyService struct {
	clientID     string
	clientSecret string
	accessToken  string
	expiresAt    time.Time
}

func NewSpotifyService() *SpotifyService {
	return &SpotifyService{
		clientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		clientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
	}
}

func (s *SpotifyService) getAccessToken() (string, error) {
	if s.accessToken != "" && s.expiresAt.After(time.Now()) {
		return s.accessToken, nil
	}

	tokenURL := "https://accounts.spotify.com/api/token"
	data := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", s.clientID, s.clientSecret)

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get access token, status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var result struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode access token response: %w", err)
	}

	s.accessToken = result.AccessToken
	s.expiresAt = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)

	return s.accessToken, nil
}

func (s *SpotifyService) SearchTracks(query string) ([]FormattedTrack, error) {
	token, err := s.getAccessToken()
	if err != nil {
		return nil, fmt.Errorf("Failed to get Spotify access token: %w", err)
	}

	LIMIT := 10
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=%d", url.QueryEscape(query), LIMIT)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Spotify search request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to make Spotify search request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Failed to search tracks, status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var searchResponse SpotifySearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		return nil, fmt.Errorf("Failed to decode Spotify search response: %w", err)
	}

	var formattedTracks []FormattedTrack
	for _, track := range searchResponse.Tracks.Items {
		artistNames := ""
		for i, artist := range track.Artists {
			if i > 0 {
				artistNames += ", "
			}
			artistNames += artist.Name
		}

		albumImageURL := "/default-image.jpg"
		if len(track.Album.Images) > 0 {
			albumImageURL = track.Album.Images[0].URL
		}

		formattedTracks = append(formattedTracks, FormattedTrack{
			SpotifyID:     track.ID,
			Name:          track.Name,
			ArtistName:    artistNames,
			AlbumImageURL: albumImageURL,
		})
	}

	return formattedTracks, nil
}
