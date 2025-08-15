package usecases

import (
	"backend/internal/infrastructure/models"
	"backend/internal/domain/repositories"
)

type PostUsecase struct {
	Repo repositories.PostRepository
}

func NewPostUsecase(repo repositories.PostRepository) *PostUsecase {
	return &PostUsecase{Repo: repo}
}

func (s PostUsecase) GetAllPosts(userID string) ([]models.Post, error) {
	return s.Repo.FindAll(userID)
}

func (s PostUsecase) GetPostByID(id uint, userID string) (*models.Post, error) {
	if userID == "" {
		return s.Repo.FindByID(id)
	}
	return s.Repo.FindByIDWithUserID(id, userID)
}

func (s PostUsecase) CreatePost(post *models.Post) error {
	return s.Repo.Create(post)
}

func (s PostUsecase) UpdatePost(post *models.Post) error {
	return s.Repo.Update(post)
}

func (s PostUsecase) DeletePost(id uint) error {
	return s.Repo.Delete(id)
}

func (s PostUsecase) GetPostsByPage(page, perPage int, userID string) ([]models.Post, int64, error) {
	return s.Repo.FindByPage(page, perPage, userID)
}