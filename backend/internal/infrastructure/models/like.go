package models

import "time"

type Like struct {
    ID        uint      `gorm:"primaryKey"`
    PostID    uint      `gorm:"not null"`
    UserID    string    `gorm:"not null"`
    CreatedAt time.Time
}
