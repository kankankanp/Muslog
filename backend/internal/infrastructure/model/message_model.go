package model

import (
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type MessageModel struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CommunityID string
	SenderID    string
	Content     string
	CreatedAt   time.Time
}

// DB → Entity
func (m *MessageModel) ToEntity() *entity.Message {
	return &entity.Message{
		ID:          m.ID,
		CommunityID: m.CommunityID,
		SenderID:    m.SenderID,
		Content:     m.Content,
		CreatedAt:   m.CreatedAt,
	}
}

// Entity → DB
func FromMessageEntity(msg *entity.Message) *MessageModel {
	return &MessageModel{
		ID:          msg.ID,
		CommunityID: msg.CommunityID,
		SenderID:    msg.SenderID,
		Content:     msg.Content,
		CreatedAt:   msg.CreatedAt,
	}
}
