package response

import "github.com/kankankanp/Muslog/internal/domain/entity"

// Spotify のトラックレスポンス
type TrackResponse struct {
	ID            uint   `json:"id"`
	SpotifyID     string `json:"spotifyId"`
	Name          string `json:"name"`
	ArtistName    string `json:"artistName"`
	AlbumImageUrl string `json:"albumImageUrl"`
	PostID        uint   `json:"postId,omitempty"`
}

func ToTrackResponse(t *entity.Track) TrackResponse {
	return TrackResponse{
		ID:            t.ID,
		SpotifyID:     t.SpotifyID,
		Name:          t.Name,
		ArtistName:    t.ArtistName,
		AlbumImageUrl: t.AlbumImageUrl,
		PostID:        t.PostID,
	}
}

func ToTrackResponses(tracks []*entity.Track) []TrackResponse {
	res := make([]TrackResponse, 0, len(tracks))
	for _, t := range tracks {
		res = append(res, ToTrackResponse(t))
	}
	return res
}

// 検索結果レスポンス
type SearchTracksResponse struct {
	Message string          `json:"message"`
	Tracks  []TrackResponse `json:"tracks"`
}
