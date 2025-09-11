package entity

import "time"

type Message struct {
	ID          string
	CommunityID string
	SenderID    string
	Content     string
	CreatedAt   time.Time
}
