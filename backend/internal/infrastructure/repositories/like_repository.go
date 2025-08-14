package repositories

import (
	"backend/internal/infrastructure/models"
	"gorm.io/gorm"
)

type likeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *likeRepository {
	return &likeRepository{DB: db}
}

func (r *likeRepository) CreateLike(like *models.Like) error {
	return r.DB.Create(like).Error
}

func (r *likeRepository) DeleteLike(postID uint, userID string) error {
	return r.DB.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&models.Like{}).Error
}

func (r *likeRepository) GetLike(postID uint, userID string) (*models.Like, error) {
	var like models.Like
	if err := r.DB.Where("post_id = ? AND user_id = ?", postID, userID).First(&like).Error; err != nil {
		return nil, err
	}
	return &like, nil
}

func (r *likeRepository) GetLikesCountByPostID(postID uint) (int, error) {
	var count int64
	if err := r.DB.Model(&models.Like{}).Where("post_id = ?", postID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}