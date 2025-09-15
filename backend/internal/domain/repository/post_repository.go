package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type PostRepository interface {
	Create(ctx context.Context, post *entity.Post) error
	FindAll(ctx context.Context, userID string) ([]entity.Post, error)
	FindByPage(ctx context.Context, page, perPage int, userID string) ([]entity.Post, int64, error)
	FindByID(ctx context.Context, id uint) (*entity.Post, error)
	FindByIDWithUserID(ctx context.Context, id uint, userID string) (*entity.Post, error)
	Update(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, id uint) error
	SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]entity.Post, int64, error)
}
