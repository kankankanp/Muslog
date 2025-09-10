package usecase

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
)

type PostUsecase interface {
	GetAllPosts(ctx context.Context, userID string) ([]entity.Post, error)
	GetPostByID(ctx context.Context, id uint, userID string) (*entity.Post, error)
	CreatePost(ctx context.Context, post *entity.Post) error
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id uint) error
	GetPostsByPage(ctx context.Context, page, perPage int, userID string) ([]entity.Post, int64, error)
	SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]entity.Post, int64, error)
}

type postUsecaseImpl struct {
	repo domainRepo.PostRepository
}

func NewPostUsecase(repo domainRepo.PostRepository) PostUsecase {
	return &postUsecaseImpl{repo: repo}
}

func (u *postUsecaseImpl) GetAllPosts(ctx context.Context, userID string) ([]entity.Post, error) {
	return u.repo.FindAll(ctx, userID)
}

func (u *postUsecaseImpl) GetPostByID(ctx context.Context, id uint, userID string) (*entity.Post, error) {
	return u.repo.FindByIDWithUserID(ctx, id, userID)
}

func (u *postUsecaseImpl) CreatePost(ctx context.Context, post *entity.Post) error {
	return u.repo.Create(ctx, post)
}

func (u *postUsecaseImpl) UpdatePost(ctx context.Context, post *entity.Post) error {
	return u.repo.Update(ctx, post)
}

func (u *postUsecaseImpl) DeletePost(ctx context.Context, id uint) error {
	return u.repo.Delete(ctx, id)
}

func (u *postUsecaseImpl) GetPostsByPage(ctx context.Context, page, perPage int, userID string) ([]entity.Post, int64, error) {
	return u.repo.FindByPage(ctx, page, perPage, userID)
}

func (u *postUsecaseImpl) SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]entity.Post, int64, error) {
	return u.repo.SearchPosts(ctx, query, tags, page, perPage, userID)
}
