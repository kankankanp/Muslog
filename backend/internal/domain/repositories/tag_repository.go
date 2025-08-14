package repositories

import "backend/internal/domain/entities"

type TagRepository interface {
	CreateTag(tag *entities.Tag) error
	GetTagByID(id uint) (*entities.Tag, error)
	GetTagByName(name string) (*entities.Tag, error)
	GetAllTags() ([]entities.Tag, error)
	UpdateTag(tag *entities.Tag) error
	DeleteTag(id uint) error
	AddTagsToPost(postID uint, tagIDs []uint) error
	RemoveTagsFromPost(postID uint, tagIDs []uint) error
	GetTagsByPostID(postID uint) ([]entities.Tag, error)
}
