package repository

import "github.com/kankankanp/Muslog/internal/domain/entity"

type TagRepository interface {
	CreateTag(tag *entity.Tag) error
	GetTagByID(id uint) (*entity.Tag, error)
	GetTagByName(name string) (*entity.Tag, error)
	GetAllTags() ([]entity.Tag, error)
	UpdateTag(tag *entity.Tag) error
	DeleteTag(id uint) error
	AddTagsToPost(postID uint, tagNames []string) error
	RemoveTagsFromPost(postID uint, tagIDs []uint) error
	GetTagsByPostID(postID uint) ([]entity.Tag, error)
}
