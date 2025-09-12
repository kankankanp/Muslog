package request

// CreatePost 用リクエスト
type CreatePostRequest struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Tracks      []TrackInput  `json:"tracks"`
}

// UpdatePost 用リクエスト
type UpdatePostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// トラック入力
type TrackInput struct {
	SpotifyID     string `json:"spotifyId"`
	Name          string `json:"name"`
	ArtistName    string `json:"artistName"`
	AlbumImageUrl string `json:"albumImageUrl"`
}