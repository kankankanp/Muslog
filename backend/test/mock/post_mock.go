package mock

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) Create(ctx context.Context, post *entity.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *MockPostRepository) FindAll(ctx context.Context, userID string) ([]*entity.Post, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Post), args.Error(1)
}

func (m *MockPostRepository) FindByPage(ctx context.Context, page, perPage int, userID string) ([]*entity.Post, int64, error) {
	args := m.Called(ctx, page, perPage, userID)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]*entity.Post), args.Get(1).(int64), args.Error(2)
}

func (m *MockPostRepository) FindByID(ctx context.Context, id uint) (*entity.Post, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Post), args.Error(1)
}

func (m *MockPostRepository) FindByIDWithUserID(ctx context.Context, id uint, userID string) (*entity.Post, error) {
	args := m.Called(ctx, id, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Post), args.Error(1)
}

func (m *MockPostRepository) Update(ctx context.Context, post *entity.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *MockPostRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPostRepository) SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]*entity.Post, int64, error) {
	args := m.Called(ctx, query, tags, page, perPage, userID)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]*entity.Post), args.Get(1).(int64), args.Error(2)
}
