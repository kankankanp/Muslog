package repositories

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"gorm.io/gorm"
)

type likeRepository struct {
	gormDB *gorm.DB
}

func NewLikeRepository(gormDB *gorm.DB) repositories.LikeRepository {
	return &likeRepository{gormDB: gormDB}
}

func (r *likeRepository) CreateLike(like *entities.Like) error {
	return r.gormDB.Create(like).Error
}

func (r *likeRepository) DeleteLike(postID uint, userID string) error {
	return r.gormDB.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&entities.Like{}).Error
}

func (r *likeRepository) GetLike(postID uint, userID string) (*entities.Like, error) {
	var like entities.Like
	err := r.gormDB.Where("post_id = ? AND user_id = ?", postID, userID).First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &like, err
}

func (r *likeRepository) GetLikesCountByPostID(postID uint) (int, error) {
	var count int64
	err := r.gormDB.Model(&entities.Like{}).Where("post_id = ?", postID).Count(&count).Error
	return int(count), err
}
