package models

type Track struct {
    ID            uint   `gorm:"primaryKey" json:"id"`
    SpotifyID     string `json:"spotifyId"`
    Name          string `json:"name"`
    ArtistName    string `json:"artistName"`
    AlbumImageUrl string `json:"albumImageUrl"`
    PostID        uint   `json:"postId"`
} 