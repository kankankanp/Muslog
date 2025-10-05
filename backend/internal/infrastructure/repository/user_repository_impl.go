package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"github.com/kankankanp/Muslog/internal/infrastructure/mapper"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domainRepo.UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var m model.UserModel
	err := r.DB.WithContext(ctx).Where("email = ?", email).First(&m).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(&m), nil
}

func (r *userRepositoryImpl) FindAll(ctx context.Context) ([]*entity.User, error) {
	var models []model.UserModel
	err := r.DB.WithContext(ctx).Find(&models).Error
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, 0, len(models))
	for _, m := range models {
		users = append(users, mapper.ToUserEntity(&m))
	}
	return users, nil
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var m model.UserModel
	err := r.DB.WithContext(ctx).
		Preload("Posts", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Preload("Posts.Tracks").
		Preload("Posts.Tags").
		First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(&m), nil
}

func (r *userRepositoryImpl) FindPosts(ctx context.Context, userID string) ([]*entity.Post, error) {
	var models []model.PostModel
	err := r.DB.WithContext(ctx).
		Where("user_id = ?", userID).
		Preload("Tracks").
		Preload("Tags").
		Order("created_at desc").
		Find(&models).Error
	if err != nil {
		return nil, err
	}

	posts := make([]*entity.Post, 0, len(models))
	for _, m := range models {
		posts = append(posts, mapper.ToPostEntity(&m))
	}
	return posts, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	m := mapper.FromUserEntity(user)
	err := r.DB.WithContext(ctx).Create(m).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToUserEntity(m), nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	m := mapper.FromUserEntity(user)
	return r.DB.WithContext(ctx).Save(m).Error
}
