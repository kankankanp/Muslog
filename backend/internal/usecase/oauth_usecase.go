package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type OAuthUsecase struct {
	UserRepo domainRepo.UserRepository
	config   *oauth2.Config
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

func NewOAuthUsecase(userRepo domainRepo.UserRepository) *OAuthUsecase {
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

	return &OAuthUsecase{
		UserRepo: userRepo,
		config:   config,
	}
}

func (s *OAuthUsecase) GetAuthURL(state string) string {
	return s.config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s *OAuthUsecase) HandleCallback(code string) (*entity.User, error) {
	token, err := s.config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %v", err)
	}

	client := s.config.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	// 既存ユーザーを探す
	existingUser, err := s.UserRepo.FindByEmail(context.Background(), userInfo.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 新規ユーザーを作成
			user := &entity.User{
				Name:     userInfo.Name,
				Email:    userInfo.Email,
				GoogleID: &userInfo.ID,
				Password: "", // OAuth認証の場合はパスワードは空
			}
			createdUser, err := s.UserRepo.Create(context.Background(), user)
			if err != nil {
				return nil, err
			}
			return createdUser, nil
		}
		// その他のDBエラー
		return nil, err
	}

	// 既存ユーザーのGoogle IDを更新
	if existingUser.GoogleID == nil || *existingUser.GoogleID != userInfo.ID {
		existingUser.GoogleID = &userInfo.ID
		err := s.UserRepo.Update(context.Background(), existingUser)
		if err != nil {
			return nil, err
		}
	}
	return existingUser, nil
}
