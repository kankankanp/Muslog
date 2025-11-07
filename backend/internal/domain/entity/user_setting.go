package entity

import "time"

type UserSetting struct {
	ID         string
	UserID     string
	EditorType string // "markdown" or "wysiwyg"
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       User
}
