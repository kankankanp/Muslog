package repository

import (
	"simple-blog/backend/internal/model"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *model.Post) error
<<<<<<< HEAD
	FindAll() ([]model.Post, error)
	FindByPage(page, perPage int) ([]model.Post, int64, error)
	FindByID(id uint) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
	GetPostByID(id uint) (*model.Post, error)
	UpdatePost(post *model.Post) error
=======
	FindAll(userID string) ([]model.Post, error)
	FindByPage(page, perPage int, userID string) ([]model.Post, int64, error)
	FindByID(id uint) (*model.Post, error) // Keep for other uses if any, or remove if not needed
	FindByIDWithUserID(id uint, userID string) (*model.Post, error)
	Update(post *model.Post) error
	Delete(id uint) error
>>>>>>> develop
}

type postRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{DB: db}
}

<<<<<<< HEAD
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
=======
func (r *postRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("Tracks").Preload("Tags").First(&post, id).Error
>>>>>>> develop
	if err != nil {
		return nil, err
	}
	return &post, nil
}

<<<<<<< HEAD
=======
func (r *postRepository) FindByIDWithUserID(id uint, userID string) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("Tracks").Preload("Tags").
		Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
		Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID).
		First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) FindAll(userID string) ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.Preload("Tracks").Preload("Tags").
		Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
		Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID).
		Order("created_at desc").
		Find(&posts).Error
	return posts, err
}

>>>>>>> develop
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

<<<<<<< HEAD
func (r *postRepository) FindByPage(page, perPage int) ([]model.Post, int64, error) {
	var posts []model.Post
	var totalCount int64
	r.DB.Model(&model.Post{}).Count(&totalCount)
	err := r.DB.Preload("Tracks").Order("created_at desc").Offset((page-1)*perPage).Limit(perPage).Find(&posts).Error
=======
func (r *postRepository) FindByPage(page, perPage int, userID string) ([]model.Post, int64, error) {
	var posts []model.Post
	var totalCount int64
	r.DB.Model(&model.Post{}).Count(&totalCount)
	err := r.DB.Preload("Tracks").Preload("Tags").
		Select("posts.*, CASE WHEN likes.user_id IS NOT NULL THEN TRUE ELSE FALSE END as is_liked").
		Joins("LEFT JOIN likes ON likes.post_id = posts.id AND likes.user_id = ?", userID).
		Order("created_at desc").
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&posts).Error
>>>>>>> develop
	return posts, totalCount, err
} 