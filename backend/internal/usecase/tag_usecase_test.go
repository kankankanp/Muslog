package usecase

import (
	"errors"
	"testing"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	testmock "github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTagUsecase_CreateTag(t *testing.T) {
	tests := []struct {
		name        string
		tagName     string
		setup       func(*testmock.MockTagRepository)
		expected    *entity.Tag
		expectedErr error
	}{
		{
			name:    "正常系: 新規タグ作成",
			tagName: "new-tag",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "new-tag").Return((*entity.Tag)(nil), errors.New("not found")).Once()
				repo.On("CreateTag", mock.MatchedBy(func(tag *entity.Tag) bool {
					assert.Equal(t, "new-tag", tag.Name)
					return true
				})).Return(nil).Run(func(args mock.Arguments) {
					tag := args.Get(0).(*entity.Tag)
					tag.ID = 1
				}).Once()
			},
			expected: &entity.Tag{ID: 1, Name: "new-tag"},
		},
		{
			name:    "異常系: 既に存在",
			tagName: "dup-tag",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "dup-tag").Return(&entity.Tag{ID: 1}, nil).Once()
			},
			expectedErr: errors.New("tag with this name already exists"),
		},
		{
			name:    "異常系: 作成失敗",
			tagName: "fail-tag",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "fail-tag").Return((*entity.Tag)(nil), errors.New("not found")).Once()
				repo.On("CreateTag", mock.AnythingOfType("*entity.Tag")).Return(errors.New("create error")).Once()
			},
			expectedErr: errors.New("create error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			tag, err := usecase.CreateTag(tt.tagName)

			if tt.expectedErr != nil {
				assert.Nil(t, tag)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Name, tag.Name)
				assert.Equal(t, tt.expected.ID, tag.ID)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}

func TestTagUsecase_GetTagByID(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*testmock.MockTagRepository)
		expected    *entity.Tag
		expectedErr error
	}{
		{
			name: "正常系: 取得成功",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByID", uint(1)).Return(&entity.Tag{ID: 1}, nil).Once()
			},
			expected: &entity.Tag{ID: 1},
		},
		{
			name: "異常系: エラー",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByID", uint(1)).Return((*entity.Tag)(nil), errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			tag, err := usecase.GetTagByID(1)

			if tt.expectedErr != nil {
				assert.Nil(t, tag)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}

func TestTagUsecase_GetTagByName(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*testmock.MockTagRepository)
		expected    *entity.Tag
		expectedErr error
	}{
		{
			name: "正常系: 取得成功",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "tag").Return(&entity.Tag{Name: "tag"}, nil).Once()
			},
			expected: &entity.Tag{Name: "tag"},
		},
		{
			name: "異常系: エラー",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "tag").Return((*entity.Tag)(nil), errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			tag, err := usecase.GetTagByName("tag")

			if tt.expectedErr != nil {
				assert.Nil(t, tag)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tag)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}

func TestTagUsecase_GetAllTags(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*testmock.MockTagRepository)
		expected    []*entity.Tag
		expectedErr error
	}{
		{
			name: "正常系: 一覧取得",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetAllTags").Return([]*entity.Tag{{ID: 1}}, nil).Once()
			},
			expected: []*entity.Tag{{ID: 1}},
		},
		{
			name: "異常系: エラー",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetAllTags").Return(nil, errors.New("list error")).Once()
			},
			expectedErr: errors.New("list error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			tags, err := usecase.GetAllTags()

			if tt.expectedErr != nil {
				assert.Nil(t, tags)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tags)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}

func TestTagUsecase_UpdateTag(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*testmock.MockTagRepository)
		expected    *entity.Tag
		expectedErr error
	}{
		{
			name: "正常系: 更新成功",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByID", uint(1)).Return(&entity.Tag{ID: 1, Name: "old"}, nil).Once()
				repo.On("GetTagByName", "new").Return((*entity.Tag)(nil), errors.New("not found")).Once()
				repo.On("UpdateTag", mock.AnythingOfType("*entity.Tag")).Return(nil).Run(func(args mock.Arguments) {
					tag := args.Get(0).(*entity.Tag)
					assert.Equal(t, "new", tag.Name)
				}).Once()
			},
			expected: &entity.Tag{ID: 1, Name: "new"},
		},
		{
			name: "異常系: 取得失敗",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByID", uint(1)).Return((*entity.Tag)(nil), errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
		{
			name: "異常系: 別IDの同名タグが存在",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByID", uint(1)).Return(&entity.Tag{ID: 1, Name: "old"}, nil).Once()
				repo.On("GetTagByName", "new").Return(&entity.Tag{ID: 2, Name: "new"}, nil).Once()
			},
			expectedErr: errors.New("tag with this name already exists"),
		},
		{
			name: "異常系: 更新失敗",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByID", uint(1)).Return(&entity.Tag{ID: 1, Name: "old"}, nil).Once()
				repo.On("GetTagByName", "new").Return((*entity.Tag)(nil), errors.New("not found")).Once()
				repo.On("UpdateTag", mock.AnythingOfType("*entity.Tag")).Return(errors.New("update error")).Once()
			},
			expectedErr: errors.New("update error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			tag, err := usecase.UpdateTag(1, "new")

			if tt.expectedErr != nil {
				assert.Nil(t, tag)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.ID, tag.ID)
				assert.Equal(t, tt.expected.Name, tag.Name)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}

func TestTagUsecase_DeleteTag(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*testmock.MockTagRepository)
		expectedErr error
	}{
		{
			name: "正常系: 削除成功",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("DeleteTag", uint(1)).Return(nil).Once()
			},
		},
		{
			name: "異常系: 削除失敗",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("DeleteTag", uint(1)).Return(errors.New("delete error")).Once()
			},
			expectedErr: errors.New("delete error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			err := usecase.DeleteTag(1)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}

func TestTagUsecase_AddTagsToPost(t *testing.T) {
	tests := []struct {
		name        string
		tagNames    []string
		setup       func(*testmock.MockTagRepository)
		expectedErr error
	}{
		{
			name:     "正常系: 既存タグのみ",
			tagNames: []string{"tag1"},
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "tag1").Return(&entity.Tag{ID: 1, Name: "tag1"}, nil).Once()
				repo.On("AddTagsToPost", uint(10), []string{"tag1"}).Return(nil).Once()
			},
		},
		{
			name:     "正常系: 未存在タグは作成",
			tagNames: []string{"tag1", "tag2"},
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "tag1").Return(&entity.Tag{ID: 1, Name: "tag1"}, nil).Once()
				repo.On("GetTagByName", "tag2").Return((*entity.Tag)(nil), errors.New("not found")).Once()
				repo.On("CreateTag", mock.AnythingOfType("*entity.Tag")).Return(nil).Run(func(args mock.Arguments) {
					tag := args.Get(0).(*entity.Tag)
					tag.ID = 2
				}).Once()
				repo.On("AddTagsToPost", uint(10), []string{"tag1", "tag2"}).Return(nil).Once()
			},
		},
		{
			name:     "異常系: タグ作成で失敗",
			tagNames: []string{"tag1"},
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "tag1").Return((*entity.Tag)(nil), errors.New("not found")).Once()
				repo.On("CreateTag", mock.AnythingOfType("*entity.Tag")).Return(errors.New("create error")).Once()
			},
			expectedErr: errors.New("create error"),
		},
		{
			name:     "異常系: 投稿への関連付けで失敗",
			tagNames: []string{"tag1"},
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "tag1").Return(&entity.Tag{ID: 1, Name: "tag1"}, nil).Once()
				repo.On("AddTagsToPost", uint(10), []string{"tag1"}).Return(errors.New("add error")).Once()
			},
			expectedErr: errors.New("add error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			err := usecase.AddTagsToPost(10, tt.tagNames)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}

func TestTagUsecase_RemoveTagsFromPost(t *testing.T) {
	tests := []struct {
		name        string
		tagNames    []string
		setup       func(*testmock.MockTagRepository)
		expectedArg []uint
		expectedErr error
	}{
		{
			name:     "正常系: タグ削除",
			tagNames: []string{"tag1", "tag2"},
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "tag1").Return(&entity.Tag{ID: 1}, nil).Once()
				repo.On("GetTagByName", "tag2").Return((*entity.Tag)(nil), errors.New("not found")).Once()
				repo.On("RemoveTagsFromPost", uint(10), []uint{1}).Return(nil).Once()
			},
			expectedArg: []uint{1},
		},
		{
			name:     "異常系: 削除処理でエラー",
			tagNames: []string{"tag1"},
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagByName", "tag1").Return(&entity.Tag{ID: 1}, nil).Once()
				repo.On("RemoveTagsFromPost", uint(10), []uint{1}).Return(errors.New("remove error")).Once()
			},
			expectedArg: []uint{1},
			expectedErr: errors.New("remove error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			err := usecase.RemoveTagsFromPost(10, tt.tagNames)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}

func TestTagUsecase_GetTagsByPostID(t *testing.T) {
	tests := []struct {
		name        string
		setup       func(*testmock.MockTagRepository)
		expected    []*entity.Tag
		expectedErr error
	}{
		{
			name: "正常系: 取得成功",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagsByPostID", uint(1)).Return([]*entity.Tag{{ID: 1}}, nil).Once()
			},
			expected: []*entity.Tag{{ID: 1}},
		},
		{
			name: "異常系: エラー",
			setup: func(repo *testmock.MockTagRepository) {
				repo.On("GetTagsByPostID", uint(1)).Return(nil, errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagRepo := new(testmock.MockTagRepository)
			postRepo := new(testmock.MockPostRepository)
			tt.setup(tagRepo)

			usecase := NewTagUsecase(tagRepo, postRepo)
			tags, err := usecase.GetTagsByPostID(1)

			if tt.expectedErr != nil {
				assert.Nil(t, tags)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tags)
			}

			tagRepo.AssertExpectations(t)
		})
	}
}
