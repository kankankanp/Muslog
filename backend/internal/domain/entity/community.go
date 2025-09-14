package entity

import "time"

type Community struct {
	ID          string
	Name        string
	Description string
	CreatorID   string
	CreatedAt   time.Time
}