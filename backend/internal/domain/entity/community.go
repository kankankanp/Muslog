package entity

import "time"

// Community represents a chat community (domain entity).
type Community struct {
	ID          string
	Name        string
	Description string
	CreatorID   string
	CreatedAt   time.Time
}