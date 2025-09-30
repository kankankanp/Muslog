package repository

import "github.com/kankankanp/Muslog/internal/domain/entity"

type LikeRepository interface {
	CreateLike(like *entity.Like) error
	DeleteLike(postID uint, userID string) error
	GetLike(postID uint, userID string) (*entity.Like, error)
	GetLikesCountByPostID(postID uint) (int, error)
	GetLikedPostsByUser(userId string) ([]*entity.Post, error)
}
