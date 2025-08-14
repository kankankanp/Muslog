package repositories

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"gorm.io/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) repositories.PostRepository {
	return &postRepository{DB: db}
}

func (r *postRepository) FindByID(id uint) (*entities.Post, error) {
	var post entities.Post
	err := r.DB.Preload("Tracks").Preload("Tags").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) FindByIDWithUserID(id uint, userID string) (*entities.Post, error) {
	var post entities.Post
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

func (r *postRepository) FindAll(userID string) ([]entities.Post, error) {
	var posts []entities.Post
	query := r.DB.Preload("Tracks").Preload("Tags")

	if userID != "" {
		query = query.
			Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
			Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
	}

	err := query.Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *postRepository) Create(post *entities.Post) error {
	return r.DB.Create(post).Error
}

func (r *postRepository) Update(post *entities.Post) error {
	return r.DB.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	// 先にtracksを削除
	err := r.DB.Where("post_id = ?", id).Delete(&entities.Track{}).Error
	if err != nil {
		return err
	}
	return r.DB.Delete(&entities.Post{}, id).Error
}

func (r *postRepository) FindByPage(page, perPage int, userID string) ([]entities.Post, int64, error) {
	var posts []entities.Post
	var totalCount int64
	r.DB.Model(&entities.Post{}).Count(&totalCount)

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