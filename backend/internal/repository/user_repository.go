package repository

import (
	"context"
	model "github.com/kankankanp/Muslog/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// FindByEmail retrieves a user by email.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll retrieves all users.
func (r *UserRepository) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.DB.WithContext(ctx).Find(&users).Error
	return users, err
}

// FindByID retrieves a user by ID.
func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Order("posts.created_at DESC")
	}).Preload("Posts.Tracks").Preload("Posts.Tags").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindPosts retrieves posts by user ID.
func (r *UserRepository) FindPosts(ctx context.Context, userID string) ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.WithContext(ctx).Where("user_id = ?", userID).Preload("Tracks").Preload("Tags").Order("created_at desc").Find(&posts).Error
	return posts, err
}

// Create creates a new user.
func (r *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.DB.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates a user's information.
func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	err := r.DB.WithContext(ctx).Save(user).Error
	if err != nil {
		return err
	}
	return nil
}
