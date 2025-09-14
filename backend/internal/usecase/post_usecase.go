package usecase

import (
	"context"
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	"github.com/kankankanp/Muslog/internal/domain/repository"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
)

type TrackInput struct {
	SpotifyID     string
	Name          string
	ArtistName    string
	AlbumImageUrl string
}

type CreatePostInput struct {
	Title          string
	Description    string
	UserID         string
	HeaderImageUrl string
	Tracks         []TrackInput
	Tags           []string
}

type PostUsecase interface {
	GetAllPosts(ctx context.Context, userID string) ([]entity.Post, error)
	GetPostByID(ctx context.Context, id uint, userID string) (*entity.Post, error)
	CreatePost(ctx context.Context, input CreatePostInput) (*entity.Post, error)
	UpdatePost(ctx context.Context, post *entity.Post) error
	DeletePost(ctx context.Context, id uint) error
	GetPostsByPage(ctx context.Context, page, perPage int, userID string) ([]entity.Post, int64, error)
	SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]entity.Post, int64, error)
}

type postUsecaseImpl struct {
	repo      domainRepo.PostRepository
	txManager repository.TransactionManager
}

func NewPostUsecase(repo domainRepo.PostRepository) PostUsecase {
	return &postUsecaseImpl{repo: repo}
}

func (u *postUsecaseImpl) GetAllPosts(ctx context.Context, userID string) ([]entity.Post, error) {
	return u.repo.FindAll(ctx, userID)
}

func (u *postUsecaseImpl) GetPostByID(ctx context.Context, id uint, userID string) (*entity.Post, error) {
	return u.repo.FindByIDWithUserID(ctx, id, userID)
}

func (u *postUsecaseImpl) CreatePost(ctx context.Context, input CreatePostInput) (*entity.Post, error) {
	var createdPost *entity.Post

	err := u.txManager.Do(ctx, func(repo repository.RepositoryProvider) error {
		post := &entity.Post{
			Title:          input.Title,
			Description:    input.Description,
			UserID:         input.UserID,
			HeaderImageUrl: input.HeaderImageUrl,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}

		if err := u.repo.Create(ctx, post); err != nil {
			return err
		}

		if len(input.Tags) > 0 {
			if err := repo.TagRepository().AddTagsToPost(post.ID, input.Tags); err != nil {
				return err
			}
		}

		createdPost = post
		return nil
	})

	if err != nil {
		return nil, err
	}
	return createdPost, nil
}

func (u *postUsecaseImpl) UpdatePost(ctx context.Context, post *entity.Post) error {
	return u.repo.Update(ctx, post)
}

func (u *postUsecaseImpl) DeletePost(ctx context.Context, id uint) error {
	return u.repo.Delete(ctx, id)
}

func (u *postUsecaseImpl) GetPostsByPage(ctx context.Context, page, perPage int, userID string) ([]entity.Post, int64, error) {
	return u.repo.FindByPage(ctx, page, perPage, userID)
}

func (u *postUsecaseImpl) SearchPosts(ctx context.Context, query string, tags []string, page, perPage int, userID string) ([]entity.Post, int64, error) {
	return u.repo.SearchPosts(ctx, query, tags, page, perPage, userID)
}
