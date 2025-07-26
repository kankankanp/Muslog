package model

import "time"

type Post struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string
    Description string
    UserID      string
    Tracks      []Track
    CreatedAt   time.Time
    UpdatedAt   time.Time
}