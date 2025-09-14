package model

import (
	"time"
)

type PostModel struct {
    ID             uint `gorm:"primaryKey"`
    Title          string
    Description    string
    UserID         string `gorm:"type:uuid"`
    HeaderImageUrl string
    Tracks         []TrackModel `gorm:"foreignKey:PostID"`
    Tags           []TagModel   `gorm:"many2many:post_tags;"`
    Likes          []LikeModel  `gorm:"foreignKey:PostID"`
    LikesCount     int          `gorm:"default:0"`
    IsLiked        bool         `gorm:"-"`
    CreatedAt      time.Time
    UpdatedAt      time.Time
}
