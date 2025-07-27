package model

import "time"

type Post struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string
    Description string
    UserID      string
    Tracks      []Track
    Tags        []Tag `gorm:"many2many:post_tags;"`
    LikesCount  int       `gorm:"default:0"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}