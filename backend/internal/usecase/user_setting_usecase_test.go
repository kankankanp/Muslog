package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

func TestUserSettingUsecase_GetUserSetting(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		mockSetting    *entity.UserSetting
		mockError      error
		expectedError  bool
		expectedResult *entity.UserSetting
	}{
		{
			name:   "設定が存在する場合",
			userID: "user-1",
			mockSetting: &entity.UserSetting{
				ID:         "setting-1",
				UserID:     "user-1",
				EditorType: "wysiwyg",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			mockError:      nil,
			expectedError:  false,
			expectedResult: &entity.UserSetting{
				ID:         "setting-1",
				UserID:     "user-1",
				EditorType: "wysiwyg",
			},
		},
		{
			name:           "設定が存在しない場合（デフォルト作成）",
			userID:         "user-2",
			mockSetting:    nil,
			mockError:      assert.AnError,
			expectedError:  false,
			expectedResult: &entity.UserSetting{
				UserID:     "user-2",
				EditorType: "markdown",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mock.MockUserSettingRepository{}
			usecase := NewUserSettingUsecase(mockRepo)

			if tt.mockError != nil {
				mockRepo.On("FindByUserID", context.Background(), tt.userID).Return(nil, tt.mockError)
				mockRepo.On("Create", context.Background(), testifyMock.MatchedBy(func(s *entity.UserSetting) bool {
					return s.UserID == tt.userID && s.EditorType == "markdown"
				})).Return(
					&entity.UserSetting{
						ID:         "new-setting",
						UserID:     tt.userID,
						EditorType: "markdown",
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}, nil)
			} else {
				mockRepo.On("FindByUserID", context.Background(), tt.userID).Return(tt.mockSetting, nil)
			}

			result, err := usecase.GetUserSetting(context.Background(), tt.userID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult.UserID, result.UserID)
				assert.Equal(t, tt.expectedResult.EditorType, result.EditorType)
			}
		})
	}
}

func TestUserSettingUsecase_UpdateUserSetting(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		editorType     string
		mockSetting    *entity.UserSetting
		findError      error
		updateError    error
		expectedError  bool
		expectedResult *entity.UserSetting
	}{
		{
			name:       "既存設定の更新",
			userID:     "user-1",
			editorType: "wysiwyg",
			mockSetting: &entity.UserSetting{
				ID:         "setting-1",
				UserID:     "user-1",
				EditorType: "markdown",
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			},
			findError:     nil,
			updateError:   nil,
			expectedError: false,
			expectedResult: &entity.UserSetting{
				UserID:     "user-1",
				EditorType: "wysiwyg",
			},
		},
		{
			name:          "新規設定の作成",
			userID:        "user-2",
			editorType:    "markdown",
			mockSetting:   nil,
			findError:     assert.AnError,
			updateError:   nil,
			expectedError: false,
			expectedResult: &entity.UserSetting{
				UserID:     "user-2",
				EditorType: "markdown",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mock.MockUserSettingRepository{}
			usecase := NewUserSettingUsecase(mockRepo)

			if tt.findError != nil {
				mockRepo.On("FindByUserID", context.Background(), tt.userID).Return(nil, tt.findError)
				mockRepo.On("Create", context.Background(), testifyMock.MatchedBy(func(s *entity.UserSetting) bool {
					return s.UserID == tt.userID && s.EditorType == tt.editorType
				})).Return(
					&entity.UserSetting{
						ID:         "new-setting",
						UserID:     tt.userID,
						EditorType: tt.editorType,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}, nil)
			} else {
				mockRepo.On("FindByUserID", context.Background(), tt.userID).Return(tt.mockSetting, nil)
				mockRepo.On("Update", context.Background(), testifyMock.MatchedBy(func(s *entity.UserSetting) bool {
					return s.UserID == tt.userID && s.EditorType == tt.editorType
				})).Return(tt.updateError)
			}

			result, err := usecase.UpdateUserSetting(context.Background(), tt.userID, tt.editorType)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult.UserID, result.UserID)
				assert.Equal(t, tt.expectedResult.EditorType, result.EditorType)
			}
		})
	}
}