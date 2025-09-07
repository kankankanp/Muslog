package usecase

import (
	"context"
	"errors"

	model "github.com/kankankanp/Muslog/internal/entity"
	"github.com/kankankanp/Muslog/internal/repository"
	gorm "gorm.io/gorm"
)

type LikeService interface {
	LikePost(ctx context.Context, postID uint, userID string) error
	UnlikePost(ctx context.Context, postID uint, userID string) error
	IsPostLikedByUser(ctx context.Context, postID uint, userID string) (bool, error)
	ToggleLike(ctx context.Context, postID uint, userID string) (bool, error) // Returns true if liked, false if unliked
}

type likeService struct {
	likeRepository repository.LikeRepository
	postRepository repository.PostRepository
}

func NewLikeService(likeRepository repository.LikeRepository, postRepository repository.PostRepository) LikeService {
	return &likeService{likeRepository: likeRepository, postRepository: postRepository}
}

func (s *likeService) LikePost(ctx context.Context, postID uint, userID string) error {
	// Check if the post exists
	post, err := s.postRepository.FindByID(ctx, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return err
	}

	// Check if already liked
	like, err := s.likeRepository.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if like != nil {
		return errors.New("post already liked by this user")
	}

	// Create like
	newLike := &model.Like{
		PostID: postID,
		UserID: userID,
	}
	if err := s.likeRepository.CreateLike(newLike); err != nil {
		return err
	}

	// Increment likes count in post
	post.LikesCount++
	return s.postRepository.Update(ctx, post)
}

func (s *likeService) UnlikePost(ctx context.Context, postID uint, userID string) error {
	// Check if the post exists
	post, err := s.postRepository.FindByID(ctx, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("post not found")
		}
		return err
	}

	// Check if liked
	like, err := s.likeRepository.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if like == nil {
		return errors.New("post not liked by this user")
	}

	// Delete like
	if err := s.likeRepository.DeleteLike(postID, userID); err != nil {
		return err
	}

	// Decrement likes count in post
	post.LikesCount--
	return s.postRepository.Update(ctx, post)
}

func (s *likeService) ToggleLike(ctx context.Context, postID uint, userID string) (bool, error) {
	// Check if the post exists
	post, err := s.postRepository.FindByID(ctx, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New("post not found")
		}
		return false, err
	}

	// Check if already liked
	like, err := s.likeRepository.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}

	if like != nil {
		// Post is already liked, so unlike it
		if err := s.likeRepository.DeleteLike(postID, userID); err != nil {
			return false, err
		}
		post.LikesCount--
		if err := s.postRepository.Update(ctx, post); err != nil {
			return false, err
		}
		return false, nil // Unliked
	} else {
		// Post is not liked, so like it
		newLike := &model.Like{
			PostID: postID,
			UserID: userID,
		}
		if err := s.likeRepository.CreateLike(newLike); err != nil {
			return false, err
		}
		post.LikesCount++
		if err := s.postRepository.Update(ctx, post); err != nil {
			return false, err
		}
		return true, nil // Liked
	}
}

func (s *likeService) IsPostLikedByUser(ctx context.Context, postID uint, userID string) (bool, error) {
	like, err := s.likeRepository.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return like != nil, nil
}