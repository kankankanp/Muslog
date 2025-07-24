package service

import (
	"backend/internal/model"
	"backend/internal/repository"
)

type BlogService struct {
	Repo *repository.BlogRepository
}

func (s *BlogService) GetAllBlogs() ([]model.Post, error) {
	return s.Repo.FindAll()
}

func (s *BlogService) GetBlogByID(id uint) (*model.Post, error) {
	return s.Repo.FindByID(id)
}

func (s *BlogService) CreateBlog(post *model.Post) error {
	return s.Repo.Create(post)
}

func (s *BlogService) UpdateBlog(post *model.Post) error {
	return s.Repo.Update(post)
}

func (s *BlogService) DeleteBlog(id uint) error {
	return s.Repo.Delete(id)
}

func (s *BlogService) GetBlogsByPage(page, perPage int) ([]model.Post, int64, error) {
	return s.Repo.FindByPage(page, perPage)
} 