package service

import (
	"errors"
	"simple-blog/backend/internal/model"
	"simple-blog/backend/internal/repository"

	gorm "gorm.io/gorm"
)

type LikeService interface {
	LikePost(postID uint, userID string) error
	UnlikePost(postID uint, userID string) error
	IsPostLikedByUser(postID uint, userID string) (bool, error)
}

type likeService struct {
	likeRepository repository.LikeRepository
	postRepository repository.PostRepository
}

func NewLikeService(likeRepository repository.LikeRepository, postRepository repository.PostRepository) LikeService {
	return &likeService{likeRepository: likeRepository, postRepository: postRepository}
}

func (s *likeService) LikePost(postID uint, userID string) error {
	// Check if the post exists
	post, err := s.postRepository.GetPostByID(postID)
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
	return s.postRepository.UpdatePost(post)
}

func (s *likeService) UnlikePost(postID uint, userID string) error {
	// Check if the post exists
	post, err := s.postRepository.GetPostByID(postID)
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
	return s.postRepository.UpdatePost(post)
}

func (s *likeService) IsPostLikedByUser(postID uint, userID string) (bool, error) {
	like, err := s.likeRepository.GetLike(postID, userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	return like != nil, nil
}
