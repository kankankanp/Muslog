package repository

import (
	model "github.com/kankankanp/Muslog/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {
	var user model.User
	err := r.DB.Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Order("posts.created_at DESC")
	}).Preload("Posts.Tracks").Preload("Posts.Tags").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindPosts(userID string) ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.Where("user_id = ?", userID).Preload("Tracks").Preload("Tags").Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	err := r.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *model.User) (*model.User, error) {
	err := r.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
