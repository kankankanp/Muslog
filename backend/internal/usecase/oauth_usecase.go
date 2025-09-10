package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/kankankanp/Muslog/internal/adapter/dto/response"
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

// --- DTO ---

type OAuthUsecase interface {
	GetAuthURL(state string) string
	HandleCallback(code string) (*entity.User, error)
}

type oAuthUsecaseImpl struct {
	userRepo domainRepo.UserRepository
	config   *oauth2.Config
}

func NewOAuthUsecase(userRepo domainRepo.UserRepository) OAuthUsecase {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return &oAuthUsecaseImpl{
		userRepo: userRepo,
		config:   config,
	}
}

func (u *oAuthUsecaseImpl) GetAuthURL(state string) string {
	return u.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (u *oAuthUsecaseImpl) HandleCallback(code string) (*entity.User, error) {
	token, err := u.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %v", err)
	}

	client := u.config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo response.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	existingUser, err := u.userRepo.FindByEmail(context.Background(), userInfo.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user := &entity.User{
				Name:     userInfo.Name,
				Email:    userInfo.Email,
				GoogleID: &userInfo.ID,
				Password: "", // OAuth 認証の場合はパスワードなし
			}
			createdUser, err := u.userRepo.Create(context.Background(), user)
			if err != nil {
				return nil, err
			}
			return createdUser, nil
		}
		return nil, err
	}

	if existingUser.GoogleID == nil || *existingUser.GoogleID != userInfo.ID {
		existingUser.GoogleID = &userInfo.ID
		if err := u.userRepo.Update(context.Background(), existingUser); err != nil {
			return nil, err
		}
	}
	return existingUser, nil
}
