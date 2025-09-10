package repository

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"gorm.io/gorm"
)

type tagRepositoryImpl struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) domainRepo.TagRepository {
	return &tagRepositoryImpl{db: db}
}

func (r *tagRepositoryImpl) CreateTag(tag *entity.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepositoryImpl) GetTagByID(id uint) (*entity.Tag, error) {
	var tag entity.Tag
	if err := r.db.First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepositoryImpl) GetTagByName(name string) (*entity.Tag, error) {
	var tag entity.Tag
	if err := r.db.Where("name = ?", name).First(&tag).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepositoryImpl) GetAllTags() ([]entity.Tag, error) {
	var tags []entity.Tag
	if err := r.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepositoryImpl) UpdateTag(tag *entity.Tag) error {
	return r.db.Save(tag).Error
}

func (r *tagRepositoryImpl) DeleteTag(id uint) error {
	return r.db.Delete(&entity.Tag{}, id).Error
}

func (r *tagRepositoryImpl) AddTagsToPost(postID uint, tagIDs []uint) error {
	var post entity.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return err
	}

	return r.db.Model(&post).Association("Tags").Append(tagIDs)
}

func (r *tagRepositoryImpl) RemoveTagsFromPost(postID uint, tagIDs []uint) error {
	var post entity.Post
	if err := r.db.First(&post, postID).Error; err != nil {
		return err
	}

	return r.db.Model(&post).Association("Tags").Delete(tagIDs)
}

func (r *tagRepositoryImpl) GetTagsByPostID(postID uint) ([]entity.Tag, error) {
	var post entity.Post
	if err := r.db.Preload("Tags").First(&post, postID).Error; err != nil {
		return nil, err
	}
	return post.Tags, nil
}
