package repository

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"gorm.io/gorm"
)

type LikeRepository interface {
	CreateLike(like *entity.Like) error
	DeleteLike(postID uint, userID string) error
	GetLike(postID uint, userID string) (*entity.Like, error)
	GetLikesCountByPostID(postID uint) (int, error)
}

type likeRepository struct {
	gormDB *gorm.DB
}

func NewLikeRepository(gormDB *gorm.DB) LikeRepository {
	return &likeRepository{gormDB: gormDB}
}

func (r *likeRepository) CreateLike(like *entity.Like) error {
	return r.gormDB.Create(like).Error
}

func (r *likeRepository) DeleteLike(postID uint, userID string) error {
	return r.gormDB.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&entity.Like{}).Error
}

func (r *likeRepository) GetLike(postID uint, userID string) (*entity.Like, error) {
	var like entity.Like
	err := r.gormDB.Where("post_id = ? AND user_id = ?", postID, userID).First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &like, err
}

func (r *likeRepository) GetLikesCountByPostID(postID uint) (int, error) {
	var count int64
	err := r.gormDB.Model(&entity.Like{}).Where("post_id = ?", postID).Count(&count).Error
	return int(count), err
}
