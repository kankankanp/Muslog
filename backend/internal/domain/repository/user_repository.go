package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
	FindPosts(ctx context.Context, userID string) ([]*entity.Post, error)
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}
