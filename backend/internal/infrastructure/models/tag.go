package models

import "time"

type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PostTag struct {
	PostID    uint      `gorm:"primaryKey" json:"postId"`
	TagID     uint      `gorm:"primaryKey" json:"tagId"`
	CreatedAt time.Time `json:"createdAt"`
}
