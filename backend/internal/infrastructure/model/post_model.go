package model

import (
	"time"
)

type PostModel struct {
	ID             uint `gorm:"primaryKey"`
	Title          string
	Description    string
	UserID         string
	HeaderImageUrl string
	Tracks         []TrackModel `gorm:"foreignKey:PostID"`
	Tags           []TagModel   `gorm:"many2many:post_tags;"`
	LikesCount     int          `gorm:"default:0"`
	IsLiked        bool         `gorm:"-"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
