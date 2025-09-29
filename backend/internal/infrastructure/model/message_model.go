package model

import (
	"time"
)

type MessageModel struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CommunityID string `gorm:"type:uuid;not null"`
	SenderID    string `gorm:"type:uuid;not null"`
	Content     string
	CreatedAt   time.Time
	Community   CommunityModel `gorm:"foreignKey:CommunityID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sender      UserModel      `gorm:"foreignKey:SenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
