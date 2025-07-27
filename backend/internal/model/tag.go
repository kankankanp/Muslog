package model

import "time"

type Tag struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostTag struct {
	PostID    uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
	CreatedAt time.Time
}
