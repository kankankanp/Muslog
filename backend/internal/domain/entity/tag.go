package entity

import "time"

type Tag struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostTag struct {
	PostID    uint
	TagID     uint
	CreatedAt time.Time
}