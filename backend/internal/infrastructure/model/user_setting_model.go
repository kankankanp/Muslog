package model

import (
	"time"
)

type UserSettingModel struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID     string `gorm:"type:uuid;not null;unique"`
	EditorType string `gorm:"default:'markdown'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       UserModel `gorm:"foreignKey:UserID"`
}