package model

import (
	"time"
)

type UserModel struct {
	ID              string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name            string
	Email           string `gorm:"unique"`
	Password        string
	GoogleID        *string `gorm:"unique;default:null"`
	ProfileImageUrl string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Posts           []PostModel       `gorm:"foreignKey:UserID"`
	Likes           []LikeModel       `gorm:"foreignKey:UserID"`
	Setting         *UserSettingModel `gorm:"foreignKey:UserID"`
}
