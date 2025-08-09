package model

type Track struct {
<<<<<<< HEAD
    ID            uint   `gorm:"primaryKey"`
    SpotifyID     string
    Name          string
    ArtistName    string
    AlbumImageUrl string
    PostID        uint
=======
    ID            uint   `gorm:"primaryKey" json:"id"`
    SpotifyID     string `json:"spotifyId"`
    Name          string `json:"name"`
    ArtistName    string `json:"artistName"`
    AlbumImageUrl string `json:"albumImageUrl"`
    PostID        uint   `json:"postId"`
>>>>>>> develop
} 