package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository" // 追加
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domainRepo.UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	err := r.DB.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := r.DB.WithContext(ctx).
		Preload("Posts", func(db *gorm.DB) *gorm.DB {
			return db.Order("posts.created_at DESC")
		}).
		Preload("Posts.Tracks").
		Preload("Posts.Tags").
		First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindPosts(ctx context.Context, userID string) ([]entity.Post, error) {
	var posts []entity.Post
	err := r.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("Tracks").
		Preload("Tags").
		Order("created_at desc").
		Find(&posts).Error
	return posts, err
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := r.DB.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	return r.DB.WithContext(ctx).Save(user).Error
}
