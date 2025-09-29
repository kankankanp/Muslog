package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestLikeUsecase_GetLikedPostsByUser(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		setupMock   func(*mock.MockLikeRepository)
		expectedLen int
		expectedErr error
	}{
		{
			name:   "正常系: いいねした投稿が複数件返る",
			userID: "test-user-uuid",
			setupMock: func(m *mock.MockLikeRepository) {
				posts := []*entity.Post{
					{ID: 1, Title: "Post1"},
					{ID: 2, Title: "Post2"},
				}
				m.On("GetLikedPostsByUser", "test-user-uuid").Return(posts, nil)
			},
			expectedLen: 2,
			expectedErr: nil,
		},
		{
			name:   "正常系: いいねした投稿が0件",
			userID: "empty-user-uuid",
			setupMock: func(m *mock.MockLikeRepository) {
				m.On("GetLikedPostsByUser", "empty-user-uuid").Return([]*entity.Post{}, nil)
			},
			expectedLen: 0,
			expectedErr: nil,
		},
		{
			name:   "異常系: リポジトリエラー",
			userID: "error-user-uuid",
			setupMock: func(m *mock.MockLikeRepository) {
				m.On("GetLikedPostsByUser", "error-user-uuid").Return(nil, errors.New("db error"))
			},
			expectedLen: 0,
			expectedErr: errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(mock.MockLikeRepository)
			tt.setupMock(mockRepo)
			// NewLikeUsecaseの引数にpostRepoも必要なら、mock.MockPostRepositoryを渡す
			usecase := NewLikeUsecase(mockRepo, nil)

			ctx := context.Background()
			posts, err := usecase.GetLikedPostsByUser(ctx, tt.userID)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Nil(t, posts)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Len(t, posts, tt.expectedLen)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
