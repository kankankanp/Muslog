package model

import (
	"time"
)

type CommunityModel struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	CreatorID   string `gorm:"not null"`
	CreatedAt   time.Time
}
