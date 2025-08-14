package repositories

import (
	"backend/internal/domain/entities"
	"backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) repositories.TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) CreateTag(tag *entities.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) GetTagByID(id uint) (*entities.Tag, error) {
	var tag entities.Tag
	if err := r.db.First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) GetTagByName(name string) (*entities.Tag, error) {
	var tag entities.Tag
	if err := r.db.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) GetAllTags() ([]entities.Tag, error) {
	var tags []entities.Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) UpdateTag(tag *entities.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepository) DeleteTag(id uint) error {
	return r.db.Delete(&entities.Tag{}, id).Error
}

func (r *tagRepository) AddTagsToPost(postID uint, tagIDs []uint) error {
	var post entities.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return err
	}

	var tags []entities.Tag
	if err := r.db.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	return r.db.Model(&post).Association("Tags").Append(&tags)
}

func (r *tagRepository) RemoveTagsFromPost(postID uint, tagIDs []uint) error {
	var post entities.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return err
	}

	var tags []entities.Tag
	if err := r.db.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	return r.db.Model(&post).Association("Tags").Delete(&tags)
}

func (r *tagRepository) GetTagsByPostID(postID uint) ([]entities.Tag, error) {
	var post entities.Post

	if err := r.db.Preload("Tags").First(&post, postID).Error; err != nil {
		return nil, err
	}

	return post.Tags, nil
}
