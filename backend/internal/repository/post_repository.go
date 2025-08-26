package repository

import (
	model "github.com/kankankanp/Muslog/internal/entity"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	FindAll(userID string) ([]model.Post, error)
	FindByPage(page, perPage int, userID string) ([]model.Post, int64, error)
	FindByID(id uint) (*model.Post, error) // Keep for other uses if any, or remove if not needed
	FindByIDWithUserID(id uint, userID string) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
}

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{DB: db}
}

func (r *postRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("Tracks").Preload("Tags").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) FindByIDWithUserID(id uint, userID string) (*model.Post, error) {
	var post model.Post
	query := r.DB.Preload("Tracks").Preload("Tags")

	if userID != "" {
		query = query.
			Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
			Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
	}

	err := query.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) FindAll(userID string) ([]model.Post, error) {
	var posts []model.Post
	query := r.DB.Preload("Tracks").Preload("Tags")

	if userID != "" {
		query = query.
			Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
			Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
	}

	err := query.Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *postRepository) Create(post *model.Post) error {
	return r.DB.Create(post).Error
}

func (r *postRepository) Update(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	// 先にtracksを削除
	err := r.DB.Where("post_id = ?", id).Delete(&model.Track{}).Error
	if err != nil {
		return err
	}
	return r.DB.Delete(&model.Post{}, id).Error
}

func (r *postRepository) FindByPage(page, perPage int, userID string) ([]model.Post, int64, error) {
	var posts []model.Post
	var totalCount int64
	r.DB.Model(&model.Post{}).Count(&totalCount)

	query := r.DB.Preload("Tracks").Preload("Tags")

	if userID != "" {
		query = query.
			Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
			Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
	}

	err := query.Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&posts).Error
	return posts, totalCount, err
}
