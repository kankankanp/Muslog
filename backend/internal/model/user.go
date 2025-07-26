package model

type User struct {
    ID       string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
    Name     string
    Email    string `gorm:"unique"`
    Password string
    Posts    []Post
} 