package model

import "time"

type Post struct {
<<<<<<< HEAD
    ID          uint      `gorm:"primaryKey"`
    Title       string
    Description string
    UserID      string
    Tracks      []Track
    Tags        []Tag `gorm:"many2many:post_tags;"`
    LikesCount  int       `gorm:"default:0"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
=======
    ID          uint      `gorm:"primaryKey" json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    UserID      string    `json:"userId"`
    Tracks      []Track   `json:"tracks"`
    Tags        []Tag     `gorm:"many2many:post_tags;" json:"tags"`
    LikesCount  int       `gorm:"default:0" json:"likesCount"`
    IsLiked     bool      `gorm:"-" json:"isLiked"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
>>>>>>> develop
}