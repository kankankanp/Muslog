package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	testmock "github.com/kankankanp/Muslog/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCommunityUsecase_CreateCommunity(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		nameInput   string
		desc        string
		creatorID   string
		setup       func(*testmock.MockCommunityRepository)
		expectedErr error
	}{
		{
			name:      "正常系: コミュニティが作成される",
			nameInput: "コミュニティ1",
			desc:      "説明",
			creatorID: "creator-1",
			setup: func(repo *testmock.MockCommunityRepository) {
				repo.On("Save", ctx, mock.MatchedBy(func(c *entity.Community) bool {
					assert.Equal(t, "コミュニティ1", c.Name)
					assert.Equal(t, "説明", c.Description)
					assert.Equal(t, "creator-1", c.CreatorID)
					assert.NotEmpty(t, c.ID)
					assert.False(t, c.CreatedAt.IsZero())
					return true
				})).Return(nil).Once()
			},
		},
		{
			name:      "異常系: 保存時にエラー",
			nameInput: "コミュニティ2",
			desc:      "説明2",
			creatorID: "creator-2",
			setup: func(repo *testmock.MockCommunityRepository) {
				repo.On("Save", ctx, mock.AnythingOfType("*entity.Community")).Return(errors.New("save error")).Once()
			},
			expectedErr: errors.New("save error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(testmock.MockCommunityRepository)
			if tt.setup != nil {
				tt.setup(repo)
			}

			usecase := NewCommunityUsecase(repo)
			community, err := usecase.CreateCommunity(ctx, tt.nameInput, tt.desc, tt.creatorID)

			if tt.expectedErr != nil {
				assert.Nil(t, community)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.nameInput, community.Name)
				assert.Equal(t, tt.desc, community.Description)
				assert.Equal(t, tt.creatorID, community.CreatorID)
				assert.NotEmpty(t, community.ID)
			}

			repo.AssertExpectations(t)
		})
	}
}

func TestCommunityUsecase_GetAllCommunities(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		setup       func(*testmock.MockCommunityRepository)
		expected    []*entity.Community
		expectedErr error
	}{
		{
			name: "正常系: 複数件取得",
			setup: func(repo *testmock.MockCommunityRepository) {
				repo.On("FindAll", ctx).Return([]*entity.Community{
					{ID: "1"},
					{ID: "2"},
				}, nil).Once()
			},
			expected: []*entity.Community{{ID: "1"}, {ID: "2"}},
		},
		{
			name: "異常系: リポジトリエラー",
			setup: func(repo *testmock.MockCommunityRepository) {
				repo.On("FindAll", ctx).Return(nil, errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(testmock.MockCommunityRepository)
			tt.setup(repo)

			usecase := NewCommunityUsecase(repo)
			communities, err := usecase.GetAllCommunities(ctx)

			if tt.expectedErr != nil {
				assert.Nil(t, communities)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, communities)
			}

			repo.AssertExpectations(t)
		})
	}
}

func TestCommunityUsecase_GetCommunityByID(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		communityID string
		setup       func(*testmock.MockCommunityRepository)
		expected    *entity.Community
		expectedErr error
	}{
		{
			name:        "正常系: コミュニティ取得",
			communityID: "community-1",
			setup: func(repo *testmock.MockCommunityRepository) {
				repo.On("FindByID", ctx, "community-1").Return(&entity.Community{ID: "community-1"}, nil).Once()
			},
			expected: &entity.Community{ID: "community-1"},
		},
		{
			name:        "異常系: リポジトリエラー",
			communityID: "community-2",
			setup: func(repo *testmock.MockCommunityRepository) {
				repo.On("FindByID", ctx, "community-2").Return((*entity.Community)(nil), errors.New("find error")).Once()
			},
			expectedErr: errors.New("find error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(testmock.MockCommunityRepository)
			tt.setup(repo)

			usecase := NewCommunityUsecase(repo)
			community, err := usecase.GetCommunityByID(ctx, tt.communityID)

			if tt.expectedErr != nil {
				assert.Nil(t, community)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, community)
			}

			repo.AssertExpectations(t)
		})
	}
}

func TestCommunityUsecase_SearchCommunities(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name        string
		query       string
		page        int
		perPage     int
		setup       func(*testmock.MockCommunityRepository)
		expected    []*entity.Community
		expectedCnt int64
		expectedErr error
	}{
		{
			name:    "正常系: ページング検索",
			query:   "music",
			page:    1,
			perPage: 10,
			setup: func(repo *testmock.MockCommunityRepository) {
				repo.On("SearchCommunities", ctx, "music", 1, 10).Return([]*entity.Community{{ID: "1"}}, int64(5), nil).Once()
			},
			expected:    []*entity.Community{{ID: "1"}},
			expectedCnt: 5,
		},
		{
			name:    "異常系: リポジトリエラー",
			query:   "error",
			page:    2,
			perPage: 5,
			setup: func(repo *testmock.MockCommunityRepository) {
				repo.On("SearchCommunities", ctx, "error", 2, 5).Return(nil, int64(0), errors.New("search error")).Once()
			},
			expectedErr: errors.New("search error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(testmock.MockCommunityRepository)
			tt.setup(repo)

			usecase := NewCommunityUsecase(repo)
			communities, total, err := usecase.SearchCommunities(ctx, tt.query, tt.page, tt.perPage)

			if tt.expectedErr != nil {
				assert.Nil(t, communities)
				assert.Zero(t, total)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, communities)
				assert.Equal(t, tt.expectedCnt, total)
			}

			repo.AssertExpectations(t)
		})
	}
}
