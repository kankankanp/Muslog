package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	testmock "github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		nameInput   string
		email       string
		password    string
		setup       func(*testmock.MockUserRepository)
		expectedErr error
	}{
		{
			name:      "正常系: ユーザー作成",
			nameInput: "user",
			email:     "user@example.com",
			password:  "password",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("Create", ctx, mock.MatchedBy(func(u *entity.User) bool {
					assert.Equal(t, "user", u.Name)
					assert.Equal(t, "user@example.com", u.Email)
					assert.NotEmpty(t, u.Password)
					return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte("password")) == nil
				})).Return(&entity.User{ID: "1"}, nil).Once()
			},
		},
		{
			name:      "異常系: 作成失敗",
			nameInput: "user",
			email:     "user@example.com",
			password:  "password",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("Create", ctx, mock.AnythingOfType("*entity.User")).Return((*entity.User)(nil), errors.New("create error")).Once()
			},
			expectedErr: errors.New("create error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(testmock.MockUserRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(userRepo)

			usecase := NewUserUsecase(userRepo, postRepo)
			user, err := usecase.CreateUser(ctx, tt.nameInput, tt.email, tt.password)

			if tt.expectedErr != nil {
				assert.Nil(t, user)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			}

			userRepo.AssertExpectations(t)
		})
	}
}

func TestUserUsecase_AuthenticateUser(t *testing.T) {
	ctx := context.Background()
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	tests := []struct {
		name        string
		email       string
		password    string
		setup       func(*testmock.MockUserRepository)
		expectedErr error
	}{
		{
			name:     "正常系: 認証成功",
			email:    "user@example.com",
			password: "password",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindByEmail", ctx, "user@example.com").Return(&entity.User{Email: "user@example.com", Password: string(hashed)}, nil).Once()
			},
		},
		{
			name:     "異常系: ユーザー未存在",
			email:    "user@example.com",
			password: "password",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindByEmail", ctx, "user@example.com").Return((*entity.User)(nil), errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
		{
			name:     "異常系: パスワード不一致",
			email:    "user@example.com",
			password: "wrong",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindByEmail", ctx, "user@example.com").Return(&entity.User{Email: "user@example.com", Password: string(hashed)}, nil).Once()
			},
			expectedErr: errors.New("mismatch"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(testmock.MockUserRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(userRepo)

			usecase := NewUserUsecase(userRepo, postRepo)
			user, err := usecase.AuthenticateUser(ctx, tt.email, tt.password)

			if tt.expectedErr != nil {
				assert.Nil(t, user)
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
			}

			userRepo.AssertExpectations(t)
		})
	}
}

func TestUserUsecase_GetAllUsers(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockUserRepository)
		expected    []*entity.User
		expectedErr error
	}{
		{
			name: "正常系: 一覧取得",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindAll", ctx).Return([]*entity.User{{ID: "1"}}, nil).Once()
			},
			expected: []*entity.User{{ID: "1"}},
		},
		{
			name: "異常系: エラー",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindAll", ctx).Return(nil, errors.New("list error")).Once()
			},
			expectedErr: errors.New("list error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(testmock.MockUserRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(userRepo)

			usecase := NewUserUsecase(userRepo, postRepo)
			users, err := usecase.GetAllUsers(ctx)

			if tt.expectedErr != nil {
				assert.Nil(t, users)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, users)
			}

			userRepo.AssertExpectations(t)
		})
	}
}

func TestUserUsecase_GetUserByID(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockUserRepository)
		expected    *entity.User
		expectedErr error
	}{
		{
			name: "正常系: 取得成功",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindByID", ctx, "user-1").Return(&entity.User{ID: "user-1"}, nil).Once()
			},
			expected: &entity.User{ID: "user-1"},
		},
		{
			name: "異常系: エラー",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindByID", ctx, "user-1").Return((*entity.User)(nil), errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(testmock.MockUserRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(userRepo)

			usecase := NewUserUsecase(userRepo, postRepo)
			user, err := usecase.GetUserByID(ctx, "user-1")

			if tt.expectedErr != nil {
				assert.Nil(t, user)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, user)
			}

			userRepo.AssertExpectations(t)
		})
	}
}

func TestUserUsecase_GetUserPosts(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockUserRepository)
		expected    []*entity.Post
		expectedErr error
	}{
		{
			name: "正常系: 投稿取得",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindPosts", ctx, "user-1").Return([]*entity.Post{{ID: 1}}, nil).Once()
			},
			expected: []*entity.Post{{ID: 1}},
		},
		{
			name: "異常系: エラー",
			setup: func(repo *testmock.MockUserRepository) {
				repo.On("FindPosts", ctx, "user-1").Return(nil, errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepo := new(testmock.MockUserRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(userRepo)

			usecase := NewUserUsecase(userRepo, postRepo)
			posts, err := usecase.GetUserPosts(ctx, "user-1")

			if tt.expectedErr != nil {
				assert.Nil(t, posts)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, posts)
			}

			userRepo.AssertExpectations(t)
		})
	}
}
