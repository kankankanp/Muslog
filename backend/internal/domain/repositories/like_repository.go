package repositories

import "backend/internal/domain/entities"

type LikeRepository interface {
	CreateLike(like *entities.Like) error
	DeleteLike(postID uint, userID string) error
	GetLike(postID uint, userID string) (*entities.Like, error)
	GetLikesCountByPostID(postID uint) (int, error)
}
