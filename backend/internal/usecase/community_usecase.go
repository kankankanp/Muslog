package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/kankankanp/Muslog/internal/entity"
	"github.com/kankankanp/Muslog/internal/repository"
)

// CommunityUsecase defines the interface for community-related business logic.
type CommunityUsecase interface {
	CreateCommunity(name, description, creatorID string) (*entity.Community, error)
	GetAllCommunities() ([]entity.Community, error)
	GetCommunityByID(id string) (*entity.Community, error)
}

// communityUsecase implements CommunityUsecase.
type communityUsecase struct {
	repo repository.CommunityRepository
}

// NewCommunityUsecase creates a new CommunityUsecase.
func NewCommunityUsecase(repo repository.CommunityRepository) CommunityUsecase {
	return &communityUsecase{repo: repo}
}

// CreateCommunity creates a new community.
func (uc *communityUsecase) CreateCommunity(name, description, creatorID string) (*entity.Community, error) {
	community := &entity.Community{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CreatorID:   creatorID,
		CreatedAt:   time.Now(),
	}

	err := uc.repo.Save(community)
	if err != nil {
		return nil, err
	}
	return community, nil
}

// GetAllCommunities retrieves all communities.
func (uc *communityUsecase) GetAllCommunities() ([]entity.Community, error) {
	return uc.repo.FindAll()
}

// GetCommunityByID retrieves a community by its ID.
func (uc *communityUsecase) GetCommunityByID(id string) (*entity.Community, error) {
	return uc.repo.FindByID(id)
}
