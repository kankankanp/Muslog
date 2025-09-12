package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
)

type CommunityUsecase interface {
	CreateCommunity(ctx context.Context, name, description, creatorID string) (*entity.Community, error)
	GetAllCommunities(ctx context.Context) ([]entity.Community, error)
	GetCommunityByID(ctx context.Context, id string) (*entity.Community, error)
	SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error)
}

type communityUsecaseImpl struct {
	repo domainRepo.CommunityRepository
}

func NewCommunityUsecase(repo domainRepo.CommunityRepository) CommunityUsecase {
	return &communityUsecaseImpl{repo: repo}
}

func (u *communityUsecaseImpl) CreateCommunity(ctx context.Context, name, description, creatorID string) (*entity.Community, error) {
	community := &entity.Community{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CreatorID:   creatorID,
		CreatedAt:   time.Now(),
	}

	if err := u.repo.Save(ctx, community); err != nil {
		return nil, err
	}
	return community, nil
}

func (u *communityUsecaseImpl) GetAllCommunities(ctx context.Context) ([]entity.Community, error) {
	return u.repo.FindAll(ctx)
}

func (u *communityUsecaseImpl) GetCommunityByID(ctx context.Context, id string) (*entity.Community, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *communityUsecaseImpl) SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error) {
	return u.repo.SearchCommunities(ctx, query, page, perPage)
}
