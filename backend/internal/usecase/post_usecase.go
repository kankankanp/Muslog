package usecase

import (
	model "github.com/kankankanp/Muslog/internal/entity"
	"github.com/kankankanp/Muslog/internal/repository"
)

type PostService struct {
	Repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s PostService) GetAllPosts(userID string) ([]model.Post, error) {
	return s.Repo.FindAll(userID)
}

func (s PostService) GetPostByID(id uint, userID string) (*model.Post, error) {
	return s.Repo.FindByIDWithUserID(id, userID)
}

func (s PostService) CreatePost(post *model.Post) error {
	return s.Repo.Create(post)
}

func (s PostService) UpdatePost(post *model.Post) error {
	return s.Repo.Update(post)
}

func (s PostService) DeletePost(id uint) error {
	return s.Repo.Delete(id)
}

func (s PostService) GetPostsByPage(page, perPage int, userID string) ([]model.Post, int64, error) {
	return s.Repo.FindByPage(page, perPage, userID)
}
