package repository

import (
	"context"
	"fmt"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"gorm.io/gorm"
)

type postRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) domainRepo.PostRepository {
	return &postRepositoryImpl{DB: db}
}

func (r *postRepositoryImpl) FindByID(ctx context.Context, id uint) (*entity.Post, error) {
	var post entity.Post
	err := r.DB.WithContext(ctx).Preload("Tracks").Preload("Tags").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepositoryImpl) FindByIDWithUserID(ctx context.Context, id uint, userID string) (*entity.Post, error) {
	var post entity.Post
	query := r.DB.WithContext(ctx).Preload("Tracks").Preload("Tags")

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

func (r *postRepositoryImpl) FindAll(ctx context.Context, userID string) ([]entity.Post, error) {
	var posts []entity.Post
	query := r.DB.WithContext(ctx).Preload("Tracks").Preload("Tags")

	if userID != "" {
		query = query.
			Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
			Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
	}

	err := query.Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *postRepositoryImpl) Create(ctx context.Context, post *entity.Post) error {
	return r.DB.WithContext(ctx).Create(post).Error
}

func (r *postRepositoryImpl) Update(ctx context.Context, post *entity.Post) error {
	return r.DB.WithContext(ctx).Save(post).Error
}

func (r *postRepositoryImpl) Delete(ctx context.Context, id uint) error {
	// 先にtracksを削除
	err := r.DB.WithContext(ctx).Where("post_id = ?", id).Delete(&entity.Track{}).Error
	if err != nil {
		return err
	}
	return r.DB.WithContext(ctx).Delete(&entity.Post{}, id).Error
}

func (r *postRepositoryImpl) FindByPage(ctx context.Context, page, perPage int, userID string) ([]entity.Post, int64, error) {
	var posts []entity.Post
	var totalCount int64
	r.DB.WithContext(ctx).Model(&entity.Post{}).Count(&totalCount)

	query := r.DB.WithContext(ctx).Preload("Tracks").Preload("Tags")

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

func (r *postRepositoryImpl) SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]entity.Post, int64, error) {
	var posts []entity.Post
	var totalCount int64

	db := r.DB.WithContext(ctx).Model(&entity.Post{}).Preload("Tracks").Preload("Tags")

	if query != "" {
		searchQuery := fmt.Sprintf("%%%s%%", query)
		db = db.Where("title ILIKE ? OR description ILIKE ?", searchQuery, searchQuery)
	}

	if len(tags) > 0 {
		db = db.Joins("JOIN post_tags pt ON pt.post_id = posts.id").
			Joins("JOIN tags t ON t.id = pt.tag_id").
			Where("t.name IN (?)", tags).
			Group("posts.id")
	}

	if userID != "" {
		db = db.Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
			Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
	}

	if err := db.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Order("created_at DESC").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, totalCount, nil
}
