package repository

import (
	"github.com/kankankanp/Muslog/internal/entity"
	"gorm.io/gorm"
)

// CommunityRepository defines the interface for community data operations.
type CommunityRepository interface {
	Save(community *entity.Community) error
	FindAll() ([]entity.Community, error)
	FindByID(id string) (*entity.Community, error)
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
func (r *communityRepository) Save(community *entity.Community) error {
	return r.DB.Create(community).Error
}

// FindAll retrieves all communities from the database.
func (r *communityRepository) FindAll() ([]entity.Community, error) {
	var communities []entity.Community
	err := r.DB.Order("created_at DESC").Find(&communities).Error
	return communities, err
}

// FindByID retrieves a community by its ID.
func (r *communityRepository) FindByID(id string) (*entity.Community, error) {
	var community entity.Community
	err := r.DB.Where("id = ?", id).First(&community).Error
	return &community, err
}
