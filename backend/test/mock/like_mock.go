package mock

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockLikeRepository struct {
	mock.Mock
}

func (m *MockLikeRepository) GetLikedPostsByUser(userID string) ([]*entity.Post, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Post), args.Error(1)
}

func (m *MockLikeRepository) CreateLike(like *entity.Like) error {
	args := m.Called(like)
	return args.Error(0)
}

func (m *MockLikeRepository) GetLike(postID uint, userID string) (*entity.Like, error) {
	args := m.Called(postID, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Like), args.Error(1)
}

func (m *MockLikeRepository) DeleteLike(postID uint, userID string) error {
	args := m.Called(postID, userID)
	return args.Error(0)
}

func (m *MockLikeRepository) GetLikesCountByPostID(postID uint) (int, error) {
	args := m.Called(postID)
	return args.Int(0), args.Error(1)
}
