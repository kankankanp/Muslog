package service

import (
	"backend/internal/model"
	"backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) AuthenticateUser(email, password string) (*model.User, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserService) GetUserPosts(userID string) ([]model.Post, error) {
	return s.Repo.FindPosts(userID)
}
