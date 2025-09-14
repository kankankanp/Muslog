package db

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/repository"
	infraRepo "github.com/kankankanp/Muslog/internal/infrastructure/repository"
	"gorm.io/gorm"
)

type gormTxManager struct {
	db *gorm.DB
}

func NewGormTxManager(db *gorm.DB) repository.TransactionManager {
	return &gormTxManager{db: db}
}

func (m *gormTxManager) Do(ctx context.Context, fn func(repo repository.RepositoryProvider) error) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		repo := &gormRepositoryProvider{db: tx}
		return fn(repo)
	})
}

type gormRepositoryProvider struct {
	db *gorm.DB
}

func (p *gormRepositoryProvider) PostRepository() repository.PostRepository {
	return infraRepo.NewPostRepository(p.db)
}

func (p *gormRepositoryProvider) TagRepository() repository.TagRepository {
	return infraRepo.NewTagRepository(p.db)
}

func (p *gormRepositoryProvider) LikeRepository() repository.LikeRepository {
	return infraRepo.NewLikeRepository(p.db)
}

func (p *gormRepositoryProvider) CommunityRepository() repository.CommunityRepository {
	return infraRepo.NewCommunityRepository(p.db)
}

func (p *gormRepositoryProvider) UserRepository() repository.UserRepository {
	return infraRepo.NewUserRepository(p.db)
}

func (p *gormRepositoryProvider) MessageRepository() repository.MessageRepository {
	return infraRepo.NewMessageRepository(p.db)
}
