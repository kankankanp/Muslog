package usecase

import (
	"errors"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
)

type TagUsecase interface {
	CreateTag(name string) (*entity.Tag, error)
	GetTagByID(id uint) (*entity.Tag, error)
	GetTagByName(name string) (*entity.Tag, error)
	GetAllTags() ([]entity.Tag, error)
	UpdateTag(id uint, name string) (*entity.Tag, error)
	DeleteTag(id uint) error
	AddTagsToPost(postID uint, tagNames []string) error
	RemoveTagsFromPost(postID uint, tagNames []string) error
	GetTagsByPostID(postID uint) ([]entity.Tag, error)
}

type tagUsecaseImpl struct {
	tagRepo  domainRepo.TagRepository
	postRepo domainRepo.PostRepository
}

func NewTagUsecase(tagRepo domainRepo.TagRepository, postRepo domainRepo.PostRepository) TagUsecase {
	return &tagUsecaseImpl{
		tagRepo:  tagRepo,
		postRepo: postRepo,
	}
}

func (u *tagUsecaseImpl) CreateTag(name string) (*entity.Tag, error) {
	if _, err := u.tagRepo.GetTagByName(name); err == nil {
		return nil, errors.New("tag with this name already exists")
	}

	tag := &entity.Tag{Name: name}
	if err := u.tagRepo.CreateTag(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (u *tagUsecaseImpl) GetTagByID(id uint) (*entity.Tag, error) {
	return u.tagRepo.GetTagByID(id)
}

func (u *tagUsecaseImpl) GetTagByName(name string) (*entity.Tag, error) {
	return u.tagRepo.GetTagByName(name)
}

func (u *tagUsecaseImpl) GetAllTags() ([]entity.Tag, error) {
	return u.tagRepo.GetAllTags()
}

func (u *tagUsecaseImpl) UpdateTag(id uint, name string) (*entity.Tag, error) {
	tag, err := u.tagRepo.GetTagByID(id)
	if err != nil {
		return nil, err
	}

	if existingTag, err := u.tagRepo.GetTagByName(name); err == nil && existingTag.ID != tag.ID {
		return nil, errors.New("tag with this name already exists")
	}

	tag.Name = name
	if err := u.tagRepo.UpdateTag(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (u *tagUsecaseImpl) DeleteTag(id uint) error {
	return u.tagRepo.DeleteTag(id)
}

func (u *tagUsecaseImpl) AddTagsToPost(postID uint, tagNames []string) error {
	var tagIDs []uint
	for _, tagName := range tagNames {
		tag, err := u.tagRepo.GetTagByName(tagName)
		if err != nil {
			newTag := &entity.Tag{Name: tagName}
			if err := u.tagRepo.CreateTag(newTag); err != nil {
				return err
			}
			tagIDs = append(tagIDs, newTag.ID)
		} else {
			tagIDs = append(tagIDs, tag.ID)
		}
	}
	return u.tagRepo.AddTagsToPost(postID, tagNames)
}

func (u *tagUsecaseImpl) RemoveTagsFromPost(postID uint, tagNames []string) error {
	var tagIDs []uint
	for _, tagName := range tagNames {
		tag, err := u.tagRepo.GetTagByName(tagName)
		if err != nil {
			continue
		}
		tagIDs = append(tagIDs, tag.ID)
	}
	return u.tagRepo.RemoveTagsFromPost(postID, tagIDs)
}

func (u *tagUsecaseImpl) GetTagsByPostID(postID uint) ([]entity.Tag, error) {
	return u.tagRepo.GetTagsByPostID(postID)
}
