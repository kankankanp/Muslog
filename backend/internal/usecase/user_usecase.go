package usecase

import (
	"context"

	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, name, email, password string) (*entity.User, error)
	AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error)
	GetAllUsers(ctx context.Context) ([]entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserPosts(ctx context.Context, userID string) ([]entity.Post, error)
}

type userUsecaseImpl struct {
	userRepo domainRepo.UserRepository
	postRepo domainRepo.PostRepository
}

func NewUserUsecase(userRepo domainRepo.UserRepository, postRepo domainRepo.PostRepository) UserUsecase {
	return &userUsecaseImpl{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (u *userUsecaseImpl) CreateUser(ctx context.Context, name, email, password string) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecaseImpl) AuthenticateUser(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecaseImpl) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return u.userRepo.FindAll(ctx)
}

func (u *userUsecaseImpl) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	return u.userRepo.FindByID(ctx, id)
}

func (u *userUsecaseImpl) GetUserPosts(ctx context.Context, userID string) ([]entity.Post, error) {
	return u.userRepo.FindPosts(ctx, userID)
}
