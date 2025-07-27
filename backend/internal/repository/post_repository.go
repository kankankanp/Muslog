package repository

import (
	"simple-blog/backend/internal/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
	FindAll() ([]model.Post, error)
	FindByPage(page, perPage int) ([]model.Post, int64, error)
	FindByID(id uint) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
	GetPostByID(id uint) (*model.Post, error)
	UpdatePost(post *model.Post) error
}

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{DB: db}
}

func (r *postRepository) GetPostByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("Tracks").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) UpdatePost(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *postRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.Preload("Tracks").Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *postRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("Tracks").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) Create(post *model.Post) error {
	return r.DB.Create(post).Error
}

func (r *postRepository) Update(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	// 先にtracksを削除
	err := r.DB.Where("post_id = ?", id).Delete(&model.Track{}).Error
	if err != nil {
		return err
	}
	return r.DB.Delete(&model.Post{}, id).Error
}

func (r *postRepository) FindByPage(page, perPage int) ([]model.Post, int64, error) {
	var posts []model.Post
	var totalCount int64
	r.DB.Model(&model.Post{}).Count(&totalCount)
	err := r.DB.Preload("Tracks").Order("created_at desc").Offset((page-1)*perPage).Limit(perPage).Find(&posts).Error
	return posts, totalCount, err
} 