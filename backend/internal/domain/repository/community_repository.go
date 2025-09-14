package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type CommunityRepository interface {
	Save(ctx context.Context, community *entity.Community) error
	FindAll(ctx context.Context) ([]entity.Community, error)
	FindByID(ctx context.Context, id string) (*entity.Community, error)
	SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error)
}