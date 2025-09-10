package repository

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"gorm.io/gorm"
)

type likeRepositoryImpl struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) domainRepo.LikeRepository {
	return &likeRepositoryImpl{db: db}
}

func (r *likeRepositoryImpl) CreateLike(like *entity.Like) error {
	return r.db.Create(like).Error
}

func (r *likeRepositoryImpl) DeleteLike(postID uint, userID string) error {
	return r.db.Where("post_id = ? AND user_id = ?", postID, userID).
		Delete(&entity.Like{}).Error
}

func (r *likeRepositoryImpl) GetLike(postID uint, userID string) (*entity.Like, error) {
	var like entity.Like
	err := r.db.Where("post_id = ? AND user_id = ?", postID, userID).
		First(&like).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &like, err
}

func (r *likeRepositoryImpl) GetLikesCountByPostID(postID uint) (int, error) {
	var count int64
	err := r.db.Model(&entity.Like{}).
		Where("post_id = ?", postID).
		Count(&count).Error
	return int(count), err
}
