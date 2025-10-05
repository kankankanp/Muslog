package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	testmock "github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostUsecase_GetAllPosts(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockPostRepository)
		expected    []*entity.Post
		expectedErr error
	}{
		{
			name: "正常系: 投稿一覧取得",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("FindAll", ctx, "user-1").Return([]*entity.Post{{ID: 1}}, nil).Once()
			},
			expected: []*entity.Post{{ID: 1}},
		},
		{
			name: "異常系: リポジトリエラー",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("FindAll", ctx, "user-1").Return(nil, errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postRepo := new(testmock.MockPostRepository)
			txManager := &testmock.MockTransactionManager{}
			if tt.setup != nil {
				tt.setup(postRepo)
			}

			usecase := NewPostUsecase(postRepo, txManager)
			posts, err := usecase.GetAllPosts(ctx, "user-1")

			if tt.expectedErr != nil {
				assert.Nil(t, posts)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, posts)
			}

			postRepo.AssertExpectations(t)
		})
	}
}

func TestPostUsecase_GetPostByID(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockPostRepository)
		expected    *entity.Post
		expectedErr error
	}{
		{
			name: "正常系: 投稿取得",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("FindByIDWithUserID", ctx, uint(1), "user-1").Return(&entity.Post{ID: 1}, nil).Once()
			},
			expected: &entity.Post{ID: 1},
		},
		{
			name: "異常系: リポジトリエラー",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("FindByIDWithUserID", ctx, uint(1), "user-1").Return((*entity.Post)(nil), errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postRepo := new(testmock.MockPostRepository)
			txManager := &testmock.MockTransactionManager{}
			tt.setup(postRepo)

			usecase := NewPostUsecase(postRepo, txManager)
			post, err := usecase.GetPostByID(ctx, 1, "user-1")

			if tt.expectedErr != nil {
				assert.Nil(t, post)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, post)
			}

			postRepo.AssertExpectations(t)
		})
	}
}

func TestPostUsecase_CreatePost(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		input       CreatePostInput
		txErr       error
		setup       func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository)
		expectedID  uint
		expectedErr error
	}{
		{
			name: "正常系: タグありで投稿作成",
			input: CreatePostInput{
				Title:       "title",
				Description: "desc",
				UserID:      "user-1",
				Tracks: []TrackInput{{
					SpotifyID:  "sp1",
					Name:       "Track 1",
					ArtistName: "Artist",
				}},
				Tags: []string{"tag1", "tag2"},
			},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("Create", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Run(func(args mock.Arguments) {
					p := args.Get(1).(*entity.Post)
					p.ID = 100
				}).Once()
				tagRepo.On("AddTagsToPost", uint(100), []string{"tag1", "tag2"}).Return(nil).Once()
				postRepo.On("FindByID", ctx, uint(100)).Return(&entity.Post{ID: 100, Title: "title"}, nil).Once()
			},
			expectedID: 100,
		},
		{
			name: "正常系: FindByIDが失敗した場合は作成済みデータを返す",
			input: CreatePostInput{
				Title:  "only",
				UserID: "user-1",
			},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("Create", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Run(func(args mock.Arguments) {
					p := args.Get(1).(*entity.Post)
					p.ID = 101
				}).Once()
				postRepo.On("FindByID", ctx, uint(101)).Return((*entity.Post)(nil), errors.New("load error")).Once()
			},
			expectedID: 101,
		},
		{
			name:        "異常系: トランザクションが開始前に失敗",
			input:       CreatePostInput{Title: "title"},
			txErr:       errors.New("tx error"),
			setup:       func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {},
			expectedErr: errors.New("tx error"),
		},
		{
			name:  "異常系: 投稿作成に失敗",
			input: CreatePostInput{Title: "title"},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("Create", ctx, mock.AnythingOfType("*entity.Post")).Return(errors.New("create error")).Once()
			},
			expectedErr: errors.New("create error"),
		},
		{
			name:  "異常系: タグ追加に失敗",
			input: CreatePostInput{Title: "title", Tags: []string{"tag1"}},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("Create", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Run(func(args mock.Arguments) {
					p := args.Get(1).(*entity.Post)
					p.ID = 102
				}).Once()
				tagRepo.On("AddTagsToPost", uint(102), []string{"tag1"}).Return(errors.New("tag error")).Once()
			},
			expectedErr: errors.New("tag error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postRepo := new(testmock.MockPostRepository)
			tagRepo := new(testmock.MockTagRepository)
			provider := &testmock.MockRepositoryProvider{
				PostRepo: postRepo,
				TagRepo:  tagRepo,
			}
			txManager := &testmock.MockTransactionManager{Provider: provider, Err: tt.txErr}
			if tt.setup != nil {
				tt.setup(postRepo, tagRepo)
			}

			usecase := NewPostUsecase(postRepo, txManager)
			post, err := usecase.CreatePost(ctx, tt.input)

			if tt.expectedErr != nil {
				assert.Nil(t, post)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, post)
				assert.Equal(t, tt.expectedID, post.ID)
			}

			postRepo.AssertExpectations(t)
			tagRepo.AssertExpectations(t)
		})
	}
}

func TestPostUsecase_UpdatePost(t *testing.T) {
	ctx := context.Background()
	now := time.Now()

	tests := []struct {
		name        string
		input       UpdatePostInput
		setup       func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository)
		expectedID  uint
		expectedErr error
	}{
		{
			name: "正常系: タグありで更新",
			input: UpdatePostInput{
				ID:             1,
				Title:          "new title",
				Description:    "new desc",
				HeaderImageUrl: "image",
				UserID:         "user-1",
				Tracks: []TrackInput{{
					SpotifyID: "sp1",
					Name:      "Track",
				}},
				Tags: []string{"tag1"},
			},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("FindByIDWithUserID", ctx, uint(1), "user-1").Return(&entity.Post{ID: 1, UpdatedAt: now}, nil).Once()
				postRepo.On("Update", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Once()
				tagRepo.On("RemoveTagsFromPost", uint(1), mock.Anything).Return(nil).Once()
				tagRepo.On("AddTagsToPost", uint(1), []string{"tag1"}).Return(nil).Once()
				postRepo.On("FindByID", ctx, uint(1)).Return(&entity.Post{ID: 1, Title: "new title"}, nil).Once()
			},
			expectedID: 1,
		},
		{
			name:  "正常系: TagRepository.AddTagsToPostはタグ無しなら呼ばれない",
			input: UpdatePostInput{ID: 2, UserID: "user-1"},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("FindByIDWithUserID", ctx, uint(2), "user-1").Return(&entity.Post{ID: 2}, nil).Once()
				postRepo.On("Update", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Once()
				tagRepo.On("RemoveTagsFromPost", uint(2), mock.Anything).Return(nil).Once()
				postRepo.On("FindByID", ctx, uint(2)).Return(&entity.Post{ID: 2}, nil).Once()
			},
			expectedID: 2,
		},
		{
			name:  "正常系: FindByID失敗時は更新後データを返す",
			input: UpdatePostInput{ID: 3, UserID: "user-1"},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("FindByIDWithUserID", ctx, uint(3), "user-1").Return(&entity.Post{ID: 3}, nil).Once()
				postRepo.On("Update", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Once()
				tagRepo.On("RemoveTagsFromPost", uint(3), mock.Anything).Return(nil).Once()
				postRepo.On("FindByID", ctx, uint(3)).Return((*entity.Post)(nil), errors.New("load error")).Once()
			},
			expectedID: 3,
		},
		{
			name:  "異常系: 投稿取得に失敗",
			input: UpdatePostInput{ID: 4, UserID: "user-1"},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("FindByIDWithUserID", ctx, uint(4), "user-1").Return((*entity.Post)(nil), errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
		{
			name:  "異常系: 更新失敗",
			input: UpdatePostInput{ID: 5, UserID: "user-1"},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("FindByIDWithUserID", ctx, uint(5), "user-1").Return(&entity.Post{ID: 5}, nil).Once()
				postRepo.On("Update", ctx, mock.AnythingOfType("*entity.Post")).Return(errors.New("update error")).Once()
			},
			expectedErr: errors.New("update error"),
		},
		{
			name:  "異常系: タグ削除に失敗",
			input: UpdatePostInput{ID: 6, UserID: "user-1"},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("FindByIDWithUserID", ctx, uint(6), "user-1").Return(&entity.Post{ID: 6}, nil).Once()
				postRepo.On("Update", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Once()
				tagRepo.On("RemoveTagsFromPost", uint(6), mock.Anything).Return(errors.New("remove error")).Once()
			},
			expectedErr: errors.New("remove error"),
		},
		{
			name:  "異常系: タグ追加時に失敗",
			input: UpdatePostInput{ID: 7, UserID: "user-1", Tags: []string{"tag1"}},
			setup: func(postRepo *testmock.MockPostRepository, tagRepo *testmock.MockTagRepository) {
				postRepo.On("FindByIDWithUserID", ctx, uint(7), "user-1").Return(&entity.Post{ID: 7}, nil).Once()
				postRepo.On("Update", ctx, mock.AnythingOfType("*entity.Post")).Return(nil).Once()
				tagRepo.On("RemoveTagsFromPost", uint(7), mock.Anything).Return(nil).Once()
				tagRepo.On("AddTagsToPost", uint(7), []string{"tag1"}).Return(errors.New("add error")).Once()
			},
			expectedErr: errors.New("add error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postRepo := new(testmock.MockPostRepository)
			tagRepo := new(testmock.MockTagRepository)
			provider := &testmock.MockRepositoryProvider{PostRepo: postRepo, TagRepo: tagRepo}
			txManager := &testmock.MockTransactionManager{Provider: provider}
			if tt.setup != nil {
				tt.setup(postRepo, tagRepo)
			}

			usecase := NewPostUsecase(postRepo, txManager)
			post, err := usecase.UpdatePost(ctx, tt.input)

			if tt.expectedErr != nil {
				assert.Nil(t, post)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, post)
				assert.Equal(t, tt.expectedID, post.ID)
			}

			postRepo.AssertExpectations(t)
			tagRepo.AssertExpectations(t)
		})
	}
}

func TestPostUsecase_DeletePost(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockPostRepository)
		expectedErr error
	}{
		{
			name: "正常系: 削除成功",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("Delete", ctx, uint(1)).Return(nil).Once()
			},
		},
		{
			name: "異常系: 削除失敗",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("Delete", ctx, uint(1)).Return(errors.New("delete error")).Once()
			},
			expectedErr: errors.New("delete error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postRepo := new(testmock.MockPostRepository)
			txManager := &testmock.MockTransactionManager{}
			tt.setup(postRepo)

			usecase := NewPostUsecase(postRepo, txManager)
			err := usecase.DeletePost(ctx, 1)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			postRepo.AssertExpectations(t)
		})
	}
}

func TestPostUsecase_GetPostsByPage(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockPostRepository)
		expected    []*entity.Post
		expectedCnt int64
		expectedErr error
	}{
		{
			name: "正常系: ページ取得",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("FindByPage", ctx, 1, 10, "user-1").Return([]*entity.Post{{ID: 1}}, int64(5), nil).Once()
			},
			expected:    []*entity.Post{{ID: 1}},
			expectedCnt: 5,
		},
		{
			name: "異常系: リポジトリエラー",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("FindByPage", ctx, 1, 10, "user-1").Return(nil, int64(0), errors.New("page error")).Once()
			},
			expectedErr: errors.New("page error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postRepo := new(testmock.MockPostRepository)
			txManager := &testmock.MockTransactionManager{}
			tt.setup(postRepo)

			usecase := NewPostUsecase(postRepo, txManager)
			posts, total, err := usecase.GetPostsByPage(ctx, 1, 10, "user-1")

			if tt.expectedErr != nil {
				assert.Nil(t, posts)
				assert.Zero(t, total)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, posts)
				assert.Equal(t, tt.expectedCnt, total)
			}

			postRepo.AssertExpectations(t)
		})
	}
}

func TestPostUsecase_SearchPosts(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockPostRepository)
		expected    []*entity.Post
		expectedCnt int64
		expectedErr error
	}{
		{
			name: "正常系: 検索成功",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("SearchPosts", ctx, "music", []string{"tag"}, 1, 10, "user-1").Return([]*entity.Post{{ID: 1}}, int64(3), nil).Once()
			},
			expected:    []*entity.Post{{ID: 1}},
			expectedCnt: 3,
		},
		{
			name: "異常系: 検索失敗",
			setup: func(repo *testmock.MockPostRepository) {
				repo.On("SearchPosts", ctx, "music", []string{"tag"}, 1, 10, "user-1").Return(nil, int64(0), errors.New("search error")).Once()
			},
			expectedErr: errors.New("search error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postRepo := new(testmock.MockPostRepository)
			txManager := &testmock.MockTransactionManager{}
			tt.setup(postRepo)

			usecase := NewPostUsecase(postRepo, txManager)
			posts, total, err := usecase.SearchPosts(ctx, "music", []string{"tag"}, 1, 10, "user-1")

			if tt.expectedErr != nil {
				assert.Nil(t, posts)
				assert.Zero(t, total)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, posts)
				assert.Equal(t, tt.expectedCnt, total)
			}

			postRepo.AssertExpectations(t)
		})
	}
}
