package model

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	Email     string    `gorm:"unique"`
	Password  string
	GoogleID  *string   `gorm:"unique;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Posts     []Post
} 