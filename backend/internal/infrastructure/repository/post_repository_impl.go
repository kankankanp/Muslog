package repository

import (
	"context"
	"fmt"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"github.com/kankankanp/Muslog/internal/infrastructure/mapper"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"gorm.io/gorm"
)

type postRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) domainRepo.PostRepository {
	return &postRepositoryImpl{DB: db}
}

func (r *postRepositoryImpl) FindByID(ctx context.Context, id uint) (*entity.Post, error) {
	var m model.PostModel
	err := r.DB.WithContext(ctx).Preload("Tracks").Preload("Tags").First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return mapper.ToPostEntity(&m), nil
}

func (r *postRepositoryImpl) FindByIDWithUserID(ctx context.Context, id uint, userID string) (*entity.Post, error) {
	var m model.PostModel
	query := r.DB.WithContext(ctx).Preload("Tracks").Preload("Tags")

    if userID != "" {
        query = query.
            Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
            Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
    }

	if err := query.First(&m, id).Error; err != nil {
		return nil, err
	}
	return mapper.ToPostEntity(&m), nil
}

func (r *postRepositoryImpl) FindAll(ctx context.Context, userID string) ([]entity.Post, error) {
	var models []model.PostModel
	query := r.DB.WithContext(ctx).Preload("Tracks").Preload("Tags")

    if userID != "" {
        query = query.
            Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
            Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
    }

    if err := query.Order("posts.created_at desc").Find(&models).Error; err != nil {
        return nil, err
    }

	posts := make([]entity.Post, 0, len(models))
	for _, m := range models {
		posts = append(posts, *mapper.ToPostEntity(&m))
	}
	return posts, nil
}

func (r *postRepositoryImpl) Create(ctx context.Context, post *entity.Post) error {
    m := mapper.FromPostEntity(post)
    if err := r.DB.WithContext(ctx).Create(m).Error; err != nil {
        return err
    }
    // write back generated fields so callers can use ID immediately
    updated := mapper.ToPostEntity(m)
    post.ID = updated.ID
    post.CreatedAt = updated.CreatedAt
    post.UpdatedAt = updated.UpdatedAt
    return nil
}

func (r *postRepositoryImpl) Update(ctx context.Context, post *entity.Post) error {
	m := mapper.FromPostEntity(post)
	return r.DB.WithContext(ctx).Save(m).Error
}

func (r *postRepositoryImpl) Delete(ctx context.Context, id uint) error {
	// Tracks 削除
	if err := r.DB.WithContext(ctx).Where("post_id = ?", id).Delete(&model.TrackModel{}).Error; err != nil {
		return err
	}
	return r.DB.WithContext(ctx).Delete(&model.PostModel{}, id).Error
}

func (r *postRepositoryImpl) FindByPage(ctx context.Context, page, perPage int, userID string) ([]entity.Post, int64, error) {
	var models []model.PostModel
	var totalCount int64
	r.DB.WithContext(ctx).Model(&model.PostModel{}).Count(&totalCount)

	query := r.DB.WithContext(ctx).Preload("Tracks").Preload("Tags")

    if userID != "" {
        query = query.
            Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
            Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID)
    }

    if err := query.Order("posts.created_at desc").
        Offset((page - 1) * perPage).
        Limit(perPage).
        Find(&models).Error; err != nil {
        return nil, 0, err
    }

	posts := make([]entity.Post, 0, len(models))
	for _, m := range models {
		posts = append(posts, *mapper.ToPostEntity(&m))
	}

	return posts, totalCount, nil
}

func (r *postRepositoryImpl) SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]entity.Post, int64, error) {
    var models []model.PostModel
    var totalCount int64

    db := r.DB.WithContext(ctx).Model(&model.PostModel{}).Preload("Tracks").Preload("Tags")

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

    // コンテキストが既にキャンセルされていれば早期リターン
    select {
    case <-ctx.Done():
        return nil, 0, ctx.Err()
    default:
    }

    if err := db.Count(&totalCount).Error; err != nil {
        return nil, 0, err
    }

    // クエリ実行前にもキャンセルを再確認
    select {
    case <-ctx.Done():
        return nil, 0, ctx.Err()
    default:
    }

    if err := db.Order("posts.created_at DESC").
        Offset((page - 1) * perPage).
        Limit(perPage).
        Find(&models).Error; err != nil {
        return nil, 0, err
    }

	posts := make([]entity.Post, 0, len(models))
	for _, m := range models {
		posts = append(posts, *mapper.ToPostEntity(&m))
	}

	return posts, totalCount, nil
}
