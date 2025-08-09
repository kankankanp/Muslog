package model

import "time"

type Tag struct {
<<<<<<< HEAD
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostTag struct {
	PostID    uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
	CreatedAt time.Time
=======
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PostTag struct {
	PostID    uint      `gorm:"primaryKey" json:"postId"`
	TagID     uint      `gorm:"primaryKey" json:"tagId"`
	CreatedAt time.Time `json:"createdAt"`
>>>>>>> develop
}
