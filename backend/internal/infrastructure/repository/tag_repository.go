package repository

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"github.com/kankankanp/Muslog/internal/infrastructure/mapper"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"gorm.io/gorm"
)

type tagRepositoryImpl struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) domainRepo.TagRepository {
	return &tagRepositoryImpl{db: db}
}

func (r *tagRepositoryImpl) CreateTag(tag *entity.Tag) error {
    m := mapper.FromTagEntity(tag)
    if err := r.db.Create(m).Error; err != nil {
        return err
    }
    // write back generated fields
    tag.ID = m.ID
    tag.CreatedAt = m.CreatedAt
    tag.UpdatedAt = m.UpdatedAt
    return nil
}

func (r *tagRepositoryImpl) GetTagByID(id uint) (*entity.Tag, error) {
	var m model.TagModel
	if err := r.db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return mapper.ToTagEntity(&m), nil
}

func (r *tagRepositoryImpl) GetTagByName(name string) (*entity.Tag, error) {
	var m model.TagModel
	if err := r.db.Where("name = ?", name).First(&m).Error; err != nil {
		return nil, err
	}
	return mapper.ToTagEntity(&m), nil
}

func (r *tagRepositoryImpl) GetAllTags() ([]entity.Tag, error) {
	var models []model.TagModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	tags := make([]entity.Tag, 0, len(models))
	for _, m := range models {
		tags = append(tags, *mapper.ToTagEntity(&m))
	}
	return tags, nil
}

func (r *tagRepositoryImpl) UpdateTag(tag *entity.Tag) error {
	m := mapper.FromTagEntity(tag)
	return r.db.Save(m).Error
}

func (r *tagRepositoryImpl) DeleteTag(id uint) error {
	return r.db.Delete(&model.TagModel{}, id).Error
}

func (r *tagRepositoryImpl) AddTagsToPost(postID uint, tagIDs []uint) error {
	// Post を model でロード
	var post model.PostModel
	if err := r.db.First(&post, postID).Error; err != nil {
		return err
	}

	// Post と Tag の関連付け（中間テーブル post_tags）
	var tags []model.TagModel
	if err := r.db.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	return r.db.Model(&post).Association("Tags").Append(&tags)
}

func (r *tagRepositoryImpl) RemoveTagsFromPost(postID uint, tagIDs []uint) error {
	var post model.PostModel
	if err := r.db.First(&post, postID).Error; err != nil {
		return err
	}

	var tags []model.TagModel
	if err := r.db.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
		return err
	}

	return r.db.Model(&post).Association("Tags").Delete(&tags)
}

func (r *tagRepositoryImpl) GetTagsByPostID(postID uint) ([]entity.Tag, error) {
	var post model.PostModel
	if err := r.db.Preload("Tags").First(&post, postID).Error; err != nil {
		return nil, err
	}

	tags := make([]entity.Tag, 0, len(post.Tags))
	for _, m := range post.Tags {
		tags = append(tags, *mapper.ToTagEntity(&m))
	}
	return tags, nil
}
