package service

import (
	"simple-blog/backend/internal/model"
	"simple-blog/backend/internal/repository"
)

type PostService struct {
	Repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s PostService) GetAllPosts() ([]model.Post, error) {
	return s.Repo.FindAll()
}

func (s PostService) GetPostByID(id uint) (*model.Post, error) {
	return s.Repo.FindByID(id)
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

func (s PostService) GetPostsByPage(page, perPage int) ([]model.Post, int64, error) {
	return s.Repo.FindByPage(page, perPage)
}