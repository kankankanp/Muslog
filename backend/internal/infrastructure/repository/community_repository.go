package repository

import (
	"context"
	"fmt"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"gorm.io/gorm"
)

type communityRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) domainRepo.CommunityRepository {
	return &communityRepositoryImpl{DB: db}
}

func (r *communityRepositoryImpl) Save(ctx context.Context, community *entity.Community) error {
	return r.DB.WithContext(ctx).Create(community).Error
}

func (r *communityRepositoryImpl) FindAll(ctx context.Context) ([]entity.Community, error) {
	var communities []entity.Community
	err := r.DB.WithContext(ctx).
		Order("created_at DESC").
		Find(&communities).Error
	return communities, err
}

func (r *communityRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Community, error) {
	var community entity.Community
	err := r.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&community).Error
	return &community, err
}

func (r *communityRepositoryImpl) SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error) {
	var communities []entity.Community
	var totalCount int64

	db := r.DB.WithContext(ctx).Model(&entity.Community{})

	if query != "" {
		searchQuery := fmt.Sprintf("%%%s%%", query)
		db = db.Where("name ILIKE ? OR description ILIKE ?", searchQuery, searchQuery)
	}

	if err := db.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((page - 1) * perPage).
		Limit(perPage).
		Order("created_at DESC").
		Find(&communities).Error; err != nil {
		return nil, 0, err
	}

	return communities, totalCount, nil
}
