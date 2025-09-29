package repository

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"github.com/kankankanp/Muslog/internal/infrastructure/mapper"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"gorm.io/gorm"
)

type likeRepositoryImpl struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) domainRepo.LikeRepository {
	return &likeRepositoryImpl{db: db}
}

func (r *likeRepositoryImpl) CreateLike(like *entity.Like) error {
	m := mapper.FromLikeEntity(like)
	return r.db.Create(m).Error
}

func (r *likeRepositoryImpl) DeleteLike(postID uint, userID string) error {
	return r.db.Where("post_id = ? AND user_id = ?", postID, userID).
		Delete(&model.LikeModel{}).Error
}

func (r *likeRepositoryImpl) GetLike(postID uint, userID string) (*entity.Like, error) {
	var m model.LikeModel
	err := r.db.Where("post_id = ? AND user_id = ?", postID, userID).
		First(&m).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mapper.ToLikeEntity(&m), err
}

func (r *likeRepositoryImpl) GetLikesCountByPostID(postID uint) (int, error) {
	var count int64
	err := r.db.Model(&model.LikeModel{}).
		Where("post_id = ?", postID).
		Count(&count).Error
	return int(count), err
}

/*
ログインユーザーがいいねした投稿一覧を取得
*/
func (r *likeRepositoryImpl) GetLikedPostsByUser(userID string) ([]*entity.Post, error) {
	var posts []*model.PostModel
	err := r.db.
		Joins("JOIN likes ON likes.post_id = post.id").
		Where("likes.user_id = ?", userID).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}

	var result []*entity.Post
	for _, m := range posts {
		result = append(result, mapper.ToPostEntity(m))
	}
	return result, nil
}
