package mock

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) Save(message *entity.Message) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MockMessageRepository) FindByCommunityID(communityID string) ([]*entity.Message, error) {
	args := m.Called(communityID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Message), args.Error(1)
}
