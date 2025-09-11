package external

// ========== Google OAuth ==========

type GoogleUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// ========== Spotify API ==========

type SpotifySearchResponse struct {
	Tracks struct {
		Items []SpotifyTrack `json:"items"`
	} `json:"tracks"`
}

type SpotifyTrack struct {
	ID     string               `json:"id"`
	Name   string               `json:"name"`
	Album  SpotifyAlbum         `json:"album"`
	Artists []SpotifyArtist     `json:"artists"`
}

type SpotifyAlbum struct {
	Images []SpotifyImage `json:"images"`
}

type SpotifyArtist struct {
	Name string `json:"name"`
}

type SpotifyImage struct {
	URL string `json:"url"`
}