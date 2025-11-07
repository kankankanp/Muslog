package entity

import "time"

type User struct {
	ID              string
	Name            string
	Email           string
	Password        string
	GoogleID        *string
	ProfileImageUrl string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Posts           []Post
	Setting         *UserSetting
}
