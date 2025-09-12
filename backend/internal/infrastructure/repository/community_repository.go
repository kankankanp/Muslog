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

type communityRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) domainRepo.CommunityRepository {
	return &communityRepositoryImpl{DB: db}
}

func (r *communityRepositoryImpl) Save(ctx context.Context, community *entity.Community) error {
	m := mapper.FromCommunityEntity(community)
	return r.DB.WithContext(ctx).Create(m).Error
}

func (r *communityRepositoryImpl) FindAll(ctx context.Context) ([]entity.Community, error) {
	var models []model.CommunityModel
	if err := r.DB.WithContext(ctx).
		Order("created_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	communities := make([]entity.Community, 0, len(models))
	for _, m := range models {
		communities = append(communities, *mapper.ToCommunityEntity(&m))
	}
	return communities, nil
}

func (r *communityRepositoryImpl) FindByID(ctx context.Context, id string) (*entity.Community, error) {
	var m model.CommunityModel
	if err := r.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error; err != nil {
		return nil, err
	}
	return mapper.ToCommunityEntity(&m), nil
}

func (r *communityRepositoryImpl) SearchCommunities(ctx context.Context, query string, page, perPage int) ([]entity.Community, int64, error) {
	var models []model.CommunityModel
	var totalCount int64

	db := r.DB.WithContext(ctx).Model(&model.CommunityModel{})

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
		Find(&models).Error; err != nil {
		return nil, 0, err
	}

	communities := make([]entity.Community, 0, len(models))
	for _, m := range models {
		communities = append(communities, *mapper.ToCommunityEntity(&m))
	}

	return communities, totalCount, nil
}
