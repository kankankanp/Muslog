package usecases

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	Repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{Repo: repo}
}

func (s *UserUsecase) CreateUser(name string, email string, password string) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.Repo.Create(user)
}

func (s *UserUsecase) AuthenticateUser(email, password string) (*entities.User, error) {
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

func (s *UserUsecase) GetAllUsers() ([]entities.User, error) {
	return s.Repo.FindAll()
}

func (s *UserUsecase) GetUserByID(id string) (*entities.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserUsecase) GetUserPosts(userID string) ([]entities.Post, error) {
	return s.Repo.FindPosts(userID)
}
