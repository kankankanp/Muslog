package repositories

import "backend/internal/infrastructure/models"

type LikeRepository interface {
	CreateLike(like *models.Like) error
	DeleteLike(postID uint, userID string) error
	GetLike(postID uint, userID string) (*models.Like, error)
	GetLikesCountByPostID(postID uint) (int, error)
}
