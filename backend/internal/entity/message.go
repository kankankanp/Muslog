package entity

import "time"

// Message represents a chat message within a community.
type Message struct {
	ID          string    `json:"id"`
	CommunityID string    `json:"communityId"`
	SenderID    string    `json:"senderId"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}
