package usecase

import (
	"context"
	"errors"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"gorm.io/gorm"
)

type LikeUsecase interface {
	LikePost(ctx context.Context, postID uint, userID string) error
	UnlikePost(ctx context.Context, postID uint, userID string) error
	IsPostLikedByUser(ctx context.Context, postID uint, userID string) (bool, error)
	ToggleLike(ctx context.Context, postID uint, userID string) (bool, error) // true if liked, false if unliked
}

type likeUsecaseImpl struct {
	likeRepo domainRepo.LikeRepository
	postRepo domainRepo.PostRepository
}

func NewLikeUsecase(likeRepo domainRepo.LikeRepository, postRepo domainRepo.PostRepository) LikeUsecase {
	return &likeUsecaseImpl{
		likeRepo: likeRepo,
		postRepo: postRepo,
	}
}

func (u *likeUsecaseImpl) LikePost(ctx context.Context, postID uint, userID string) error {
	post, err := u.postRepo.FindByID(ctx, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return err
	}

	like, err := u.likeRepo.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if like != nil {
		return errors.New("post already liked by this user")
	}

	newLike := &entity.Like{PostID: postID, UserID: userID}
	if err := u.likeRepo.CreateLike(newLike); err != nil {
		return err
	}

	post.LikesCount++
	return u.postRepo.Update(ctx, post)
}

func (u *likeUsecaseImpl) UnlikePost(ctx context.Context, postID uint, userID string) error {
	post, err := u.postRepo.FindByID(ctx, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return err
	}

	like, err := u.likeRepo.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if like == nil {
		return errors.New("post not liked by this user")
	}

	if err := u.likeRepo.DeleteLike(postID, userID); err != nil {
		return err
	}

	post.LikesCount--
	return u.postRepo.Update(ctx, post)
}

func (u *likeUsecaseImpl) ToggleLike(ctx context.Context, postID uint, userID string) (bool, error) {
	post, err := u.postRepo.FindByID(ctx, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("post not found")
		}
		return false, err
	}

	like, err := u.likeRepo.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}

	if like != nil {
		if err := u.likeRepo.DeleteLike(postID, userID); err != nil {
			return false, err
		}
		post.LikesCount--
		if err := u.postRepo.Update(ctx, post); err != nil {
			return false, err
		}
		return false, nil
	}

	newLike := &entity.Like{PostID: postID, UserID: userID}
	if err := u.likeRepo.CreateLike(newLike); err != nil {
		return false, err
	}
	post.LikesCount++
	if err := u.postRepo.Update(ctx, post); err != nil {
		return false, err
	}
	return true, nil
}

func (u *likeUsecaseImpl) IsPostLikedByUser(ctx context.Context, postID uint, userID string) (bool, error) {
	like, err := u.likeRepo.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return like != nil, nil
}
