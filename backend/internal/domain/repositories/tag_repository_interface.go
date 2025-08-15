package repositories

import "backend/internal/infrastructure/models"

type TagRepository interface {
	CreateTag(tag *models.Tag) error
	GetTagByID(id uint) (*models.Tag, error)
	GetTagByName(name string) (*models.Tag, error)
	GetAllTags() ([]models.Tag, error)
	UpdateTag(tag *models.Tag) error
	DeleteTag(id uint) error
	AddTagsToPost(postID uint, tagIDs []uint) error
	RemoveTagsFromPost(postID uint, tagIDs []uint) error
	GetTagsByPostID(postID uint) ([]models.Tag, error)
}
