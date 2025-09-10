package repository

import (
	"context"
	"fmt"
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"gorm.io/gorm"
)

// CommunityRepository defines the interface for community data operations.
type CommunityRepository interface {
	Save(ctx context.Context, community *entity.Community) error
	FindAll(ctx context.Context) ([]entity.Community, error)
	FindByID(ctx context.Context, id string) (*entity.Community, error)
	SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error)
}

// communityRepository implements CommunityRepository using GORM.
type communityRepository struct {
	DB *gorm.DB
}

// NewCommunityRepository creates a new CommunityRepository.
func NewCommunityRepository(db *gorm.DB) CommunityRepository {
	return &communityRepository{DB: db}
}

// Save saves a community to the database.
func (r *communityRepository) Save(ctx context.Context, community *entity.Community) error {
	return r.DB.WithContext(ctx).Create(community).Error
}

// FindAll retrieves all communities from the database.
func (r *communityRepository) FindAll(ctx context.Context) ([]entity.Community, error) {
	var communities []entity.Community
	err := r.DB.WithContext(ctx).Order("created_at DESC").Find(&communities).Error
	return communities, err
}

// FindByID retrieves a community by its ID.
func (r *communityRepository) FindByID(ctx context.Context, id string) (*entity.Community, error) {
	var community entity.Community
	err := r.DB.WithContext(ctx).Where("id = ?", id).First(&community).Error
	return &community, err
}

// SearchCommunities searches for communities by name or description.
func (r *communityRepository) SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error) {
	var communities []entity.Community
	var totalCount int64

	db := r.DB.WithContext(ctx).Model(&entity.Community{})

	if query != "" {
		searchQuery := fmt.Sprintf("%%%s%%", query)
		db = db.Where("name ILIKE ? OR description ILIKE ?", searchQuery, searchQuery)
	}

	err := db.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * perPage).Limit(perPage).Order("created_at DESC").Find(&communities).Error
	if err != nil {
		return nil, 0, err
	}

	return communities, totalCount, nil
}