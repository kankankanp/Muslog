package usecase

import (
	"context"

	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type PostUsecase struct {
	Repo domainRepo.PostRepository
}

func NewPostUsecase(repo domainRepo.PostRepository) *PostUsecase {
	return &PostUsecase{Repo: repo}
}

func (s PostUsecase) GetAllPosts(ctx context.Context, userID string) ([]entity.Post, error) {
	return s.Repo.FindAll(ctx, userID)
}

func (s PostUsecase) GetPostByID(ctx context.Context, id uint, userID string) (*entity.Post, error) {
	return s.Repo.FindByIDWithUserID(ctx, id, userID)
}

func (s PostUsecase) CreatePost(ctx context.Context, post *entity.Post) error {
	return s.Repo.Create(ctx, post)
}

func (s PostUsecase) UpdatePost(ctx context.Context, post *entity.Post) error {
	return s.Repo.Update(ctx, post)
}

func (s PostUsecase) DeletePost(ctx context.Context, id uint) error {
	return s.Repo.Delete(ctx, id)
}

func (s PostUsecase) GetPostsByPage(ctx context.Context, page, perPage int, userID string) ([]entity.Post, int64, error) {
	return s.Repo.FindByPage(ctx, page, perPage, userID)
}

// SearchPosts searches for posts by query and tags.
func (s PostUsecase) SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]entity.Post, int64, error) {
	return s.Repo.SearchPosts(ctx, query, tags, page, perPage, userID)
}
