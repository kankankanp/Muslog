package mock

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockTagRepository struct {
	mock.Mock
}

func (m *MockTagRepository) CreateTag(tag *entity.Tag) error {
	args := m.Called(tag)
	return args.Error(0)
}

func (m *MockTagRepository) GetTagByID(id uint) (*entity.Tag, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Tag), args.Error(1)
}

func (m *MockTagRepository) GetTagByName(name string) (*entity.Tag, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Tag), args.Error(1)
}

func (m *MockTagRepository) GetAllTags() ([]*entity.Tag, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Tag), args.Error(1)
}

func (m *MockTagRepository) UpdateTag(tag *entity.Tag) error {
	args := m.Called(tag)
	return args.Error(0)
}

func (m *MockTagRepository) DeleteTag(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTagRepository) AddTagsToPost(postID uint, tagNames []string) error {
	args := m.Called(postID, tagNames)
	return args.Error(0)
}

func (m *MockTagRepository) RemoveTagsFromPost(postID uint, tagIDs []uint) error {
	args := m.Called(postID, tagIDs)
	return args.Error(0)
}

func (m *MockTagRepository) GetTagsByPostID(postID uint) ([]*entity.Tag, error) {
	args := m.Called(postID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Tag), args.Error(1)
}
