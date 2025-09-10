package usecase

import (
	"context"
	model "github.com/kankankanp/Muslog/internal/entity"
	"github.com/kankankanp/Muslog/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	Repo *repository.UserRepository
}

func (s *UserUsecase) CreateUser(ctx context.Context, name string, email string, password string) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.Repo.Create(ctx, user)
}

func (s *UserUsecase) AuthenticateUser(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.Repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserUsecase) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.Repo.FindAll(ctx)
}

func (s *UserUsecase) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *UserUsecase) GetUserPosts(ctx context.Context, userID string) ([]model.Post, error) {
	return s.Repo.FindPosts(ctx, userID)
}