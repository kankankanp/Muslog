package model

import (
	"time"
)

type LikeModel struct {
	ID        uint   `gorm:"primaryKey"`
	PostID    uint   `gorm:"not null"`
	UserID    string `gorm:"type:uuid;not null"`
	CreatedAt time.Time
	Post      PostModel `gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      UserModel `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
