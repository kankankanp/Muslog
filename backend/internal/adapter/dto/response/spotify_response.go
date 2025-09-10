package response

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
