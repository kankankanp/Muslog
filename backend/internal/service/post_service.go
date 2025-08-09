package service

import (
	"simple-blog/backend/internal/model"
	"simple-blog/backend/internal/repository"
)

<<<<<<< HEAD
type PostService interface {
	GetAllPosts() ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	CreatePost(post *model.Post) error
	UpdatePost(post *model.Post) error
	DeletePost(id uint) error
	GetPostsByPage(page, perPage int) ([]model.Post, int64, error)
}

type postService struct {
	Repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{Repo: repo}
}

func (s *postService) GetAllPosts() ([]model.Post, error) {
	return s.Repo.FindAll()
}

func (s *postService) GetPostByID(id uint) (*model.Post, error) {
	return s.Repo.FindByID(id)
}

func (s *postService) CreatePost(post *model.Post) error {
	return s.Repo.Create(post)
}

func (s *postService) UpdatePost(post *model.Post) error {
	return s.Repo.Update(post)
}

func (s *postService) DeletePost(id uint) error {
	return s.Repo.Delete(id)
}

func (s *postService) GetPostsByPage(page, perPage int) ([]model.Post, int64, error) {
	return s.Repo.FindByPage(page, perPage)
=======
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
>>>>>>> develop
}