package repositories

import "backend/internal/domain/entities"

type UserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	FindAll() ([]entities.User, error)
	FindByID(id string) (*entities.User, error)
	FindPosts(userID string) ([]entities.Post, error)
	Create(user *entities.User) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
}
