package model

import (
	"time"
)

type CommunityModel struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	CreatorID   string `gorm:"type:uuid;not null"`
	CreatedAt   time.Time
	Creator     UserModel      `gorm:"foreignKey:CreatorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Messages    []MessageModel `gorm:"foreignKey:CommunityID"`
}
