package usecase

import (
	"errors"
	"testing"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	testmock "github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
)

func TestMessageUsecase_SaveMessage(t *testing.T) {
	tests := []struct {
		name        string
		message     *entity.Message
		setup       func(*testmock.MockMessageRepository)
		expectedErr error
	}{
		{
			name:    "正常系: メッセージ保存成功",
			message: &entity.Message{ID: "1"},
			setup: func(repo *testmock.MockMessageRepository) {
				repo.On("Save", &entity.Message{ID: "1"}).Return(nil).Once()
			},
		},
		{
			name:    "異常系: 保存失敗",
			message: &entity.Message{ID: "2"},
			setup: func(repo *testmock.MockMessageRepository) {
				repo.On("Save", &entity.Message{ID: "2"}).Return(errors.New("save error")).Once()
			},
			expectedErr: errors.New("save error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(testmock.MockMessageRepository)
			tt.setup(repo)

			usecase := NewMessageUsecase(repo)
			err := usecase.SaveMessage(tt.message)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			repo.AssertExpectations(t)
		})
	}
}

func TestMessageUsecase_GetMessagesByCommunityID(t *testing.T) {
	tests := []struct {
		name        string
		communityID string
		setup       func(*testmock.MockMessageRepository)
		expected    []*entity.Message
		expectedErr error
	}{
		{
			name:        "正常系: メッセージ取得",
			communityID: "community-1",
			setup: func(repo *testmock.MockMessageRepository) {
				repo.On("FindByCommunityID", "community-1").Return([]*entity.Message{{ID: "1"}}, nil).Once()
			},
			expected: []*entity.Message{{ID: "1"}},
		},
		{
			name:        "異常系: 取得失敗",
			communityID: "community-2",
			setup: func(repo *testmock.MockMessageRepository) {
				repo.On("FindByCommunityID", "community-2").Return(nil, errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(testmock.MockMessageRepository)
			tt.setup(repo)

			usecase := NewMessageUsecase(repo)
			messages, err := usecase.GetMessagesByCommunityID(tt.communityID)

			if tt.expectedErr != nil {
				assert.Nil(t, messages)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, messages)
			}

			repo.AssertExpectations(t)
		})
	}
}
