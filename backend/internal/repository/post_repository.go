package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func (r *PostRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.Preload("Tracks").Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *PostRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("Tracks").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) Create(post *model.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) Update(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *PostRepository) Delete(id uint) error {
	// 先にtracksを削除
	err := r.DB.Where("post_id = ?", id).Delete(&model.Track{}).Error
	if err != nil {
		return err
	}
	return r.DB.Delete(&model.Post{}, id).Error
}

func (r *PostRepository) FindByPage(page, perPage int) ([]model.Post, int64, error) {
	var posts []model.Post
	var totalCount int64
	r.DB.Model(&model.Post{}).Count(&totalCount)
	err := r.DB.Preload("Tracks").Order("created_at desc").Offset((page-1)*perPage).Limit(perPage).Find(&posts).Error
	return posts, totalCount, err
} 