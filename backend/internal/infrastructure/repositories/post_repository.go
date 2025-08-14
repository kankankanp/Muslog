package repositories

import (
	"backend/internal/infrastructure/models"
	"gorm.io/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{DB: db}
}

func (r *postRepository) Create(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *postRepository) FindAll(userID string) ([]models.Post, error) {
	var posts []models.Post
	query := r.DB.Order("created_at desc")
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) FindByPage(page, perPage int, userID string) ([]models.Post, int64, error) {
	var posts []models.Post
	var totalCount int64

	query := r.DB.Model(&models.Post{}).Order("created_at desc")
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	// Count total posts
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Paginate results
	offset := (page - 1) * perPage
	if err := query.Limit(perPage).Offset(offset).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, totalCount, nil
}

func (r *postRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.DB.Preload("Tracks").Preload("Tags").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) FindByIDWithUserID(id uint, userID string) (*models.Post, error) {
	var post models.Post
	if err := r.DB.Preload("Tracks").Preload("Tags").Where("id = ? AND user_id = ?", id, userID).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Update(post *models.Post) error {
	return r.DB.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Post{}, id).Error
}