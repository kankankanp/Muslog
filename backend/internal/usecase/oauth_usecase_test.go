package usecase

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	testmock "github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newJSONResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     http.Header{"Content-Type": []string{"application/json"}},
	}
}

func TestOAuthUsecase_GetAuthURL(t *testing.T) {
	usecase := &oAuthUsecaseImpl{
		config: &oauth2.Config{
			ClientID:    "client-id",
			RedirectURL: "http://localhost/callback",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://example.com/auth",
				TokenURL: "https://example.com/token",
			},
			Scopes: []string{"scope1"},
		},
	}

	url := usecase.GetAuthURL("state-token")
	assert.Contains(t, url, "state=state-token")
	assert.Contains(t, url, "scope=")
}

func TestOAuthUsecase_HandleCallback(t *testing.T) {
	tests := []struct {
		name        string
		setupRepo   func(*testmock.MockUserRepository)
		setupHTTP   func(t *testing.T) func()
		expected    *entity.User
		expectedErr error
	}{
		{
			name: "正常系: 新規ユーザー作成",
			setupRepo: func(repo *testmock.MockUserRepository) {
				repo.On("FindByEmail", mock.Anything, "user@example.com").Return((*entity.User)(nil), gorm.ErrRecordNotFound).Once()
				repo.On("Create", mock.Anything, mock.MatchedBy(func(u *entity.User) bool {
					assert.Equal(t, "User Name", u.Name)
					assert.Equal(t, "user@example.com", u.Email)
					assert.Empty(t, u.Password)
					return true
				})).Return(&entity.User{Email: "user@example.com"}, nil).Once()
			},
			setupHTTP: func(t *testing.T) func() {
				t.Helper()
				originalClient := http.DefaultClient
				http.DefaultClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					switch {
					case req.URL.Host == "oauth2.googleapis.com" && req.URL.Path == "/token":
						return newJSONResponse(http.StatusOK, `{"access_token":"token123","token_type":"Bearer","expires_in":3600}`), nil
					case req.URL.Host == "www.googleapis.com" && req.URL.Path == "/oauth2/v2/userinfo":
						return newJSONResponse(http.StatusOK, `{"id":"google-id","email":"user@example.com","name":"User Name"}`), nil
					default:
						return nil, errors.New("unexpected request: " + req.URL.String())
					}
				})}
				return func() {
					http.DefaultClient = originalClient
				}
			},
			expected: &entity.User{Email: "user@example.com"},
		},
		{
			name: "正常系: 既存ユーザー更新",
			setupRepo: func(repo *testmock.MockUserRepository) {
				existing := &entity.User{Email: "user@example.com"}
				repo.On("FindByEmail", mock.Anything, "user@example.com").Return(existing, nil).Once()
				repo.On("Update", mock.Anything, mock.MatchedBy(func(u *entity.User) bool {
					assert.Equal(t, "google-id", *u.GoogleID)
					return true
				})).Return(nil).Once()
			},
			setupHTTP: func(t *testing.T) func() {
				t.Helper()
				originalClient := http.DefaultClient
				http.DefaultClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					switch {
					case req.URL.Host == "oauth2.googleapis.com" && req.URL.Path == "/token":
						return newJSONResponse(http.StatusOK, `{"access_token":"token123","token_type":"Bearer","expires_in":3600}`), nil
					case req.URL.Host == "www.googleapis.com" && req.URL.Path == "/oauth2/v2/userinfo":
						return newJSONResponse(http.StatusOK, `{"id":"google-id","email":"user@example.com","name":"User Name"}`), nil
					default:
						return nil, errors.New("unexpected request: " + req.URL.String())
					}
				})}
				return func() {
					http.DefaultClient = originalClient
				}
			},
			expected: &entity.User{Email: "user@example.com"},
		},
		{
			name:      "異常系: トークン交換失敗",
			setupRepo: func(repo *testmock.MockUserRepository) {},
			setupHTTP: func(t *testing.T) func() {
				t.Helper()
				originalClient := http.DefaultClient
				http.DefaultClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					if req.URL.Host == "oauth2.googleapis.com" {
						return newJSONResponse(http.StatusBadRequest, `{"error":"invalid_grant"}`), nil
					}
					return nil, errors.New("unexpected request")
				})}
				return func() {
					http.DefaultClient = originalClient
				}
			},
			expectedErr: errors.New("failed to exchange code"),
		},
		{
			name:      "異常系: ユーザー情報取得失敗",
			setupRepo: func(repo *testmock.MockUserRepository) {},
			setupHTTP: func(t *testing.T) func() {
				t.Helper()
				originalClient := http.DefaultClient
				http.DefaultClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					switch {
					case req.URL.Host == "oauth2.googleapis.com" && req.URL.Path == "/token":
						return newJSONResponse(http.StatusOK, `{"access_token":"token123","token_type":"Bearer","expires_in":3600}`), nil
					case req.URL.Host == "www.googleapis.com" && req.URL.Path == "/oauth2/v2/userinfo":
						return nil, errors.New("network error")
					default:
						return nil, errors.New("unexpected request")
					}
				})}
				return func() {
					http.DefaultClient = originalClient
				}
			},
			expectedErr: errors.New("failed to get user info"),
		},
		{
			name: "異常系: ユーザー作成時のリポジトリエラー",
			setupRepo: func(repo *testmock.MockUserRepository) {
				repo.On("FindByEmail", mock.Anything, "user@example.com").Return((*entity.User)(nil), gorm.ErrRecordNotFound).Once()
				repo.On("Create", mock.Anything, mock.AnythingOfType("*entity.User")).Return((*entity.User)(nil), errors.New("create error")).Once()
			},
			setupHTTP: func(t *testing.T) func() {
				t.Helper()
				originalClient := http.DefaultClient
				http.DefaultClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					switch {
					case req.URL.Host == "oauth2.googleapis.com" && req.URL.Path == "/token":
						return newJSONResponse(http.StatusOK, `{"access_token":"token123","token_type":"Bearer","expires_in":3600}`), nil
					case req.URL.Host == "www.googleapis.com" && req.URL.Path == "/oauth2/v2/userinfo":
						return newJSONResponse(http.StatusOK, `{"id":"google-id","email":"user@example.com","name":"User Name"}`), nil
					default:
						return nil, errors.New("unexpected request")
					}
				})}
				return func() {
					http.DefaultClient = originalClient
				}
			},
			expectedErr: errors.New("create error"),
		},
		{
			name: "異常系: ユーザー更新時のリポジトリエラー",
			setupRepo: func(repo *testmock.MockUserRepository) {
				existing := &entity.User{Email: "user@example.com"}
				repo.On("FindByEmail", mock.Anything, "user@example.com").Return(existing, nil).Once()
				repo.On("Update", mock.Anything, mock.AnythingOfType("*entity.User")).Return(errors.New("update error")).Once()
			},
			setupHTTP: func(t *testing.T) func() {
				t.Helper()
				originalClient := http.DefaultClient
				http.DefaultClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					switch {
					case req.URL.Host == "oauth2.googleapis.com" && req.URL.Path == "/token":
						return newJSONResponse(http.StatusOK, `{"access_token":"token123","token_type":"Bearer","expires_in":3600}`), nil
					case req.URL.Host == "www.googleapis.com" && req.URL.Path == "/oauth2/v2/userinfo":
						return newJSONResponse(http.StatusOK, `{"id":"google-id","email":"user@example.com","name":"User Name"}`), nil
					default:
						return nil, errors.New("unexpected request")
					}
				})}
				return func() {
					http.DefaultClient = originalClient
				}
			},
			expectedErr: errors.New("update error"),
		},
		{
			name: "異常系: FindByEmailが失敗",
			setupRepo: func(repo *testmock.MockUserRepository) {
				repo.On("FindByEmail", mock.Anything, "user@example.com").Return((*entity.User)(nil), errors.New("find error")).Once()
			},
			setupHTTP: func(t *testing.T) func() {
				t.Helper()
				originalClient := http.DefaultClient
				http.DefaultClient = &http.Client{Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
					switch {
					case req.URL.Host == "oauth2.googleapis.com" && req.URL.Path == "/token":
						return newJSONResponse(http.StatusOK, `{"access_token":"token123","token_type":"Bearer","expires_in":3600}`), nil
					case req.URL.Host == "www.googleapis.com" && req.URL.Path == "/oauth2/v2/userinfo":
						return newJSONResponse(http.StatusOK, `{"id":"google-id","email":"user@example.com","name":"User Name"}`), nil
					default:
						return nil, errors.New("unexpected request")
					}
				})}
				return func() {
					http.DefaultClient = originalClient
				}
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(testmock.MockUserRepository)
			if tt.setupRepo != nil {
				tt.setupRepo(repo)
			}

			cleanup := func() {}
			if tt.setupHTTP != nil {
				cleanup = tt.setupHTTP(t)
			}
			defer cleanup()

			usecase := &oAuthUsecaseImpl{
				userRepo: repo,
				config: &oauth2.Config{
					ClientID:     "client",
					ClientSecret: "secret",
					RedirectURL:  "http://localhost/callback",
					Scopes:       []string{"email"},
					Endpoint: oauth2.Endpoint{
						AuthURL:  "https://accounts.google.com/o/oauth2/auth",
						TokenURL: "https://oauth2.googleapis.com/token",
					},
				},
			}

			user, err := usecase.HandleCallback("auth-code")

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				if tt.expected != nil {
					assert.Equal(t, tt.expected.Email, user.Email)
				}
			}

			repo.AssertExpectations(t)
		})
	}
}
