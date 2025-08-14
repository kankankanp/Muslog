package repositories

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id string) (*entities.User, error) {
	var user entities.User
	err := r.DB.Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Order("posts.created_at DESC")
	}).Preload("Posts.Tracks").Preload("Posts.Tags").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindPosts(userID string) ([]entities.Post, error) {
	var posts []entities.Post
	err := r.DB.Where("user_id = ?", userID).Preload("Tracks").Preload("Tags").Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *userRepository) Create(user *entities.User) (*entities.User, error) {
	err := r.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *entities.User) (*entities.User, error) {
	err := r.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
} 
