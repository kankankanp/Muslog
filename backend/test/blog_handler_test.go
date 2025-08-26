package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kankankanp/Muslog/internal/entity"
	"github.com/kankankanp/Muslog/internal/handler"
	"github.com/kankankanp/Muslog/internal/middleware"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPostService is a mock implementation of service.PostService
type MockPostService struct {
	mock.Mock
}

func (m *MockPostService) GetAllPosts() ([]entity.Post, error) {
	args := m.Called()
	return args.Get(0).([]entity.Post), args.Error(1)
}

func (m *MockPostService) GetPostByID(id uint) (*entity.Post, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Post), args.Error(1)
}

func (m *MockPostService) CreatePost(post *entity.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *MockPostService) UpdatePost(post *entity.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *MockPostService) DeletePost(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPostService) GetPostsByPage(page, pageSize int) ([]entity.Post, int64, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]entity.Post), args.Get(1).(int64), args.Error(2)
}

func TestGetAllPosts(t *testing.T) {
	e := echo.New()
	mockPostService := new(MockPostService)
	mockPostService.On("GetAllPosts").Return([]entity.Post{}, nil) // Mock empty posts, no error
	h := handler.NewPostHandler(mockPostService)

	req := httptest.NewRequest(http.MethodGet, "/api/posts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.GetAllPosts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestCreatePostUnauthorized(t *testing.T) {
	e := echo.New()

	// Mock the middleware to simulate an unauthorized response
	e.Use(middleware.AuthMiddleware(middleware.AuthMiddlewareConfig{
		Skipper: func(c echo.Context) bool {
			// Skip only for paths that don't require auth, for this test, we want auth to fail
			return false
		},
		JWTSecret: "test_secret", // Provide a dummy secret for the middleware to initialize
	}))

	post := map[string]string{"title": "Test Title", "content": "Test Content"}
	jsonBody, _ := json.Marshal(post)
	req := httptest.NewRequest(http.MethodPost, "/blogs", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Manually call the middleware and then the handler
	authMiddlewareFunc := middleware.AuthMiddleware(middleware.AuthMiddlewareConfig{JWTSecret: "test_secret"})

	// Create a dummy next handler that just returns nil, as the middleware should prevent reaching it
	dummyNextHandler := func(c echo.Context) error { return nil }

	err := authMiddlewareFunc(dummyNextHandler)(c)

	assert.NoError(t, err) // Middleware should not return an error directly, but set the response code
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
