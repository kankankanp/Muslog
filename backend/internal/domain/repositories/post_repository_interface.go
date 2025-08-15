package repositories

import "backend/internal/infrastructure/models"

type PostRepository interface {
	Create(post *models.Post) error
	FindAll(userID string) ([]models.Post, error)
	FindByPage(page, perPage int, userID string) ([]models.Post, int64, error)
	FindByID(id uint) (*models.Post, error)
	FindByIDWithUserID(id uint, userID string) (*models.Post, error)
	Update(post *models.Post) error
	Delete(id uint) error
}
