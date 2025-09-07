package entity

import "time"

type User struct {
	ID              string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name            string
	Email           string `gorm:"unique"`
	Password        string
	GoogleID        *string `gorm:"unique;default:null"`
	ProfileImageUrl string  // New field for profile image URL
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Posts           []Post
}
