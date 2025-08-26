package repository

import (
	model "github.com/kankankanp/Muslog/internal/entity"
	"gorm.io/gorm"
)

type TagRepository interface {
	CreateTag(tag *model.Tag) error
	GetTagByID(id uint) (*model.Tag, error)
	GetTagByName(name string) (*model.Tag, error)
	GetAllTags() ([]model.Tag, error)
	UpdateTag(tag *model.Tag) error
	DeleteTag(id uint) error
	AddTagsToPost(postID uint, tagIDs []uint) error
	RemoveTagsFromPost(postID uint, tagIDs []uint) error
	GetTagsByPostID(postID uint) ([]model.Tag, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) CreateTag(tag *model.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) GetTagByID(id uint) (*model.Tag, error) {
	var tag model.Tag
	if err := r.db.First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) GetTagByName(name string) (*model.Tag, error) {
	var tag model.Tag
	if err := r.db.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) GetAllTags() ([]model.Tag, error) {
	var tags []model.Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) UpdateTag(tag *model.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepository) DeleteTag(id uint) error {
	return r.db.Delete(&model.Tag{}, id).Error
}

func (r *tagRepository) AddTagsToPost(postID uint, tagIDs []uint) error {
	var post model.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return err
	}

	return r.db.Model(&post).Association("Tags").Append(tagIDs)
}

func (r *tagRepository) RemoveTagsFromPost(postID uint, tagIDs []uint) error {
	var post model.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return err
	}

	return r.db.Model(&post).Association("Tags").Delete(tagIDs)
}

func (r *tagRepository) GetTagsByPostID(postID uint) ([]model.Tag, error) {
	var post model.Post

	if err := r.db.Preload("Tags").First(&post, postID).Error; err != nil {
		return nil, err
	}

	return post.Tags, nil
}
