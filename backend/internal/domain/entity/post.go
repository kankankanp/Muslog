package entity

import "time"

type Post struct {
	ID             uint
	Title          string
	Description    string
	UserID         string
	HeaderImageUrl string
	Tracks         []Track
	Tags           []Tag
	LikesCount     int
	IsLiked        bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
