package entity

import "time"

type Like struct {
	ID        uint
	PostID    uint
	UserID    string
	CreatedAt time.Time
}
