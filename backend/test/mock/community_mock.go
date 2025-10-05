package mock

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockCommunityRepository struct {
	mock.Mock
}

func (m *MockCommunityRepository) Save(ctx context.Context, community *entity.Community) error {
	args := m.Called(ctx, community)
	return args.Error(0)
}

func (m *MockCommunityRepository) FindAll(ctx context.Context) ([]*entity.Community, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Community), args.Error(1)
}

func (m *MockCommunityRepository) FindByID(ctx context.Context, id string) (*entity.Community, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Community), args.Error(1)
}

func (m *MockCommunityRepository) SearchCommunities(ctx context.Context, query string, page, perPage int) ([]*entity.Community, int64, error) {
	args := m.Called(ctx, query, page, perPage)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]*entity.Community), args.Get(1).(int64), args.Error(2)
}
