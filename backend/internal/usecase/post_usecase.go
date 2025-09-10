package usecase

import (
	"context"

	model "github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/kankankanp/Muslog/internal/infrastructure/repository"
)

type PostService struct {
	Repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s PostService) GetAllPosts(ctx context.Context, userID string) ([]model.Post, error) {
	return s.Repo.FindAll(ctx, userID)
}

func (s PostService) GetPostByID(ctx context.Context, id uint, userID string) (*model.Post, error) {
	return s.Repo.FindByIDWithUserID(ctx, id, userID)
}

func (s PostService) CreatePost(ctx context.Context, post *model.Post) error {
	return s.Repo.Create(ctx, post)
}

func (s PostService) UpdatePost(ctx context.Context, post *model.Post) error {
	return s.Repo.Update(ctx, post)
}

func (s PostService) DeletePost(ctx context.Context, id uint) error {
	return s.Repo.Delete(ctx, id)
}

func (s PostService) GetPostsByPage(ctx context.Context, page, perPage int, userID string) ([]model.Post, int64, error) {
	return s.Repo.FindByPage(ctx, page, perPage, userID)
}

// SearchPosts searches for posts by query and tags.
func (s PostService) SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]model.Post, int64, error) {
	return s.Repo.SearchPosts(ctx, query, tags, page, perPage, userID)
}
