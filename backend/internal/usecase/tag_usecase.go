package usecase

import (
	"errors"

	model "github.com/kankankanp/Muslog/internal/entity"
	"github.com/kankankanp/Muslog/internal/repository"
)

type TagService interface {
	CreateTag(name string) (*model.Tag, error)
	GetTagByID(id uint) (*model.Tag, error)
	GetTagByName(name string) (*model.Tag, error)
	GetAllTags() ([]model.Tag, error)
	UpdateTag(id uint, name string) (*model.Tag, error)
	DeleteTag(id uint) error
	AddTagsToPost(postID uint, tagNames []string) error
	RemoveTagsFromPost(postID uint, tagNames []string) error
	GetTagsByPostID(postID uint) ([]model.Tag, error)
}

type tagService struct {
	tagRepo  repository.TagRepository
	postRepo repository.PostRepository
}

func NewTagService(tagRepo repository.TagRepository, postRepo repository.PostRepository) TagService {
	return &tagService{tagRepo: tagRepo, postRepo: postRepo}
}

func (s *tagService) CreateTag(name string) (*model.Tag, error) {
	// Check if tag already exists
	if _, err := s.tagRepo.GetTagByName(name); err == nil {
		return nil, errors.New("tag with this name already exists")
	}

	tag := &model.Tag{Name: name}
	if err := s.tagRepo.CreateTag(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *tagService) GetTagByID(id uint) (*model.Tag, error) {
	return s.tagRepo.GetTagByID(id)
}

func (s *tagService) GetTagByName(name string) (*model.Tag, error) {
	return s.tagRepo.GetTagByName(name)
}

func (s *tagService) GetAllTags() ([]model.Tag, error) {
	return s.tagRepo.GetAllTags()
}

func (s *tagService) UpdateTag(id uint, name string) (*model.Tag, error) {
	tag, err := s.tagRepo.GetTagByID(id)
	if err != nil {
		return nil, err
	}
	tag.Name = name
	// Check if updated tag name already exists for another tag
	if existingTag, err := s.tagRepo.GetTagByName(name); err == nil && existingTag.ID != tag.ID {
		return nil, errors.New("tag with this name already exists")
	}
	if err := s.tagRepo.UpdateTag(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *tagService) DeleteTag(id uint) error {
	return s.tagRepo.DeleteTag(id)
}

func (s *tagService) AddTagsToPost(postID uint, tagNames []string) error {
	var tagIDs []uint
	for _, tagName := range tagNames {
		tag, err := s.tagRepo.GetTagByName(tagName)
		if err != nil {
			// Tag does not exist, create it
			newTag := &model.Tag{Name: tagName}
			if err := s.tagRepo.CreateTag(newTag); err != nil {
				return err
			}
			tagIDs = append(tagIDs, newTag.ID)
		} else {
			tagIDs = append(tagIDs, tag.ID)
		}
	}
	return s.tagRepo.AddTagsToPost(postID, tagIDs)
}

func (s *tagService) RemoveTagsFromPost(postID uint, tagNames []string) error {
	var tagIDs []uint
	for _, tagName := range tagNames {
		tag, err := s.tagRepo.GetTagByName(tagName)
		if err != nil {
			// Tag not found, skip or return error based on desired behavior
			continue
		}
		tagIDs = append(tagIDs, tag.ID)
	}
	return s.tagRepo.RemoveTagsFromPost(postID, tagIDs)
}

func (s *tagService) GetTagsByPostID(postID uint) ([]model.Tag, error) {
	return s.tagRepo.GetTagsByPostID(postID)
}
