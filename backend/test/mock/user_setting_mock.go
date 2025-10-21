package mock

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserSettingRepository struct {
	mock.Mock
}

func (m *MockUserSettingRepository) FindByUserID(ctx context.Context, userID string) (*entity.UserSetting, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserSetting), args.Error(1)
}

func (m *MockUserSettingRepository) Create(ctx context.Context, setting *entity.UserSetting) (*entity.UserSetting, error) {
	args := m.Called(ctx, setting)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserSetting), args.Error(1)
}

func (m *MockUserSettingRepository) Update(ctx context.Context, setting *entity.UserSetting) error {
	args := m.Called(ctx, setting)
	return args.Error(0)
}