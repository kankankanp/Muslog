package repositories

import (
	"backend/internal/infrastructure/models"
	"gorm.io/gorm"
)

type tagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) *tagRepository {
	return &tagRepository{DB: db}
}

func (r *tagRepository) CreateTag(tag *models.Tag) error {
	return r.DB.Create(tag).Error
}

func (r *tagRepository) GetTagByID(id uint) (*models.Tag, error) {
	var tag models.Tag
	if err := r.DB.First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) GetTagByName(name string) (*models.Tag, error) {
	var tag models.Tag
	if err := r.DB.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) GetAllTags() ([]models.Tag, error) {
	var tags []models.Tag
	if err := r.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) UpdateTag(tag *models.Tag) error {
	return r.DB.Save(tag).Error
}

func (r *tagRepository) DeleteTag(id uint) error {
	return r.DB.Delete(&models.Tag{}, id).Error
}

func (r *tagRepository) AddTagsToPost(postID uint, tagIDs []uint) error {
	var post models.Post
	if err := r.DB.Preload("Tags").First(&post, postID).Error; err != nil {
		return err
	}

	var tags []models.Tag
	if err := r.DB.Where(tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	return r.DB.Model(&post).Association("Tags").Append(tags)
}

func (r *tagRepository) RemoveTagsFromPost(postID uint, tagIDs []uint) error {
	var post models.Post
	if err := r.DB.Preload("Tags").First(&post, postID).Error; err != nil {
		return err
	}

	var tags []models.Tag
	if err := r.DB.Where(tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	return r.DB.Model(&post).Association("Tags").Delete(tags)
}

func (r *tagRepository) GetTagsByPostID(postID uint) ([]models.Tag, error) {
	var post models.Post
	if err := r.DB.Preload("Tags").First(&post, postID).Error; err != nil {
		return nil, err
	}
	return post.Tags, nil
}