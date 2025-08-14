package repositories

import "backend/internal/domain/entities"

type PostRepository interface {
	Create(post *entities.Post) error
	FindAll(userID string) ([]entities.Post, error)
	FindByPage(page, perPage int, userID string) ([]entities.Post, int64, error)
	FindByID(id uint) (*entities.Post, error)
	FindByIDWithUserID(id uint, userID string) (*entities.Post, error)
	Update(post *entities.Post) error
	Delete(id uint) error
}
