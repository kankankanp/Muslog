package model

type TrackModel struct {
	ID            uint `gorm:"primaryKey"`
	SpotifyID     string
	Name          string
	ArtistName    string
	AlbumImageUrl string
	PostID        uint
}
