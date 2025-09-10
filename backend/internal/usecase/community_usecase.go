package usecase

import (
	"context"
	"time"

	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"github.com/google/uuid"
	"github.com/kankankanp/Muslog/internal/domain/entity"
)

// CommunityUsecase defines the interface for community-related business logic.
type CommunityUsecase interface {
	CreateCommunity(ctx context.Context, name, description, creatorID string) (*entity.Community, error)
	GetAllCommunities(ctx context.Context) ([]entity.Community, error)
	GetCommunityByID(ctx context.Context, id string) (*entity.Community, error)
	SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error)
}

// communityUsecase implements CommunityUsecase.
type communityUsecase struct {
	repo domainRepo.CommunityRepository
}

// NewCommunityUsecase creates a new CommunityUsecase.
func NewCommunityUsecase(repo domainRepo.CommunityRepository) CommunityUsecase {
	return &communityUsecase{repo: repo}
}

// CreateCommunity creates a new community.
func (uc *communityUsecase) CreateCommunity(ctx context.Context, name, description, creatorID string) (*entity.Community, error) {
	community := &entity.Community{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CreatorID:   creatorID,
		CreatedAt:   time.Now(),
	}

	err := uc.repo.Save(ctx, community)
	if err != nil {
		return nil, err
	}
	return community, nil
}

// GetAllCommunities retrieves all communities.
func (uc *communityUsecase) GetAllCommunities(ctx context.Context) ([]entity.Community, error) {
	return uc.repo.FindAll(ctx)
}

// GetCommunityByID retrieves a community by its ID.
func (uc *communityUsecase) GetCommunityByID(ctx context.Context, id string) (*entity.Community, error) {
	return uc.repo.FindByID(ctx, id)
}

// SearchCommunities searches for communities by query.
func (uc *communityUsecase) SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error) {
	return uc.repo.SearchCommunities(ctx, query, page, perPage)
}
