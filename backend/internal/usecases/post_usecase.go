package usecases

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
)

type PostUsecase struct {
	Repo repositories.PostRepository
}

func NewPostUsecase(repo repositories.PostRepository) *PostUsecase {
	return &PostUsecase{Repo: repo}
}

func (s PostUsecase) GetAllPosts(userID string) ([]entities.Post, error) {
	return s.Repo.FindAll(userID)
}

func (s PostUsecase) GetPostByID(id uint, userID string) (*entities.Post, error) {
	return s.Repo.FindByIDWithUserID(id, userID)
}

func (s PostUsecase) CreatePost(post *entities.Post) error {
	return s.Repo.Create(post)
}

func (s PostUsecase) UpdatePost(post *entities.Post) error {
	return s.Repo.Update(post)
}

func (s PostUsecase) DeletePost(id uint) error {
	return s.Repo.Delete(id)
}

func (s PostUsecase) GetPostsByPage(page, perPage int, userID string) ([]entities.Post, int64, error) {
	return s.Repo.FindByPage(page, perPage, userID)
}