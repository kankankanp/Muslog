package service

import (
	"backend/internal/model"
	"backend/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
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