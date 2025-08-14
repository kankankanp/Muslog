package repositories

import "backend/internal/infrastructure/models"

type UserRepository interface {
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id string) error
	FindAll() ([]models.User, error)
}