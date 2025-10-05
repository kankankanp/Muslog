package mock

import (
	"context"

	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
)

type MockRepositoryProvider struct {
	PostRepo      domainRepo.PostRepository
	TagRepo       domainRepo.TagRepository
	UserRepo      domainRepo.UserRepository
	MessageRepo   domainRepo.MessageRepository
	CommunityRepo domainRepo.CommunityRepository
	LikeRepo      domainRepo.LikeRepository
}

func (m *MockRepositoryProvider) PostRepository() domainRepo.PostRepository {
	return m.PostRepo
}

func (m *MockRepositoryProvider) TagRepository() domainRepo.TagRepository {
	return m.TagRepo
}

func (m *MockRepositoryProvider) UserRepository() domainRepo.UserRepository {
	return m.UserRepo
}

func (m *MockRepositoryProvider) MessageRepository() domainRepo.MessageRepository {
	return m.MessageRepo
}

func (m *MockRepositoryProvider) CommunityRepository() domainRepo.CommunityRepository {
	return m.CommunityRepo
}

func (m *MockRepositoryProvider) LikeRepository() domainRepo.LikeRepository {
	return m.LikeRepo
}

type MockTransactionManager struct {
	Provider domainRepo.RepositoryProvider
	Err      error
}

func (m *MockTransactionManager) Do(ctx context.Context, fn func(domainRepo.RepositoryProvider) error) error {
	if m.Err != nil {
		return m.Err
	}
	if fn == nil {
		return nil
	}
	return fn(m.Provider)
}
