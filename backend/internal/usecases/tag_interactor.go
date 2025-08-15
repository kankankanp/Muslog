package usecases

import (
	"backend/internal/infrastructure/models"
	"backend/internal/domain/repositories"
	"errors"
)

type TagUsecase interface {
	CreateTag(name string) (*models.Tag, error)
	GetTagByID(id uint) (*models.Tag, error)
	GetTagByName(name string) (*models.Tag, error)
	GetAllTags() ([]models.Tag, error)
	UpdateTag(id uint, name string) (*models.Tag, error)
	DeleteTag(id uint) error
	AddTagsToPost(postID uint, tagNames []string) error
	RemoveTagsFromPost(postID uint, tagNames []string) error
	GetTagsByPostID(postID uint) ([]models.Tag, error)
}

type tagUsecase struct {
	tagRepo  repositories.TagRepository
	postRepo repositories.PostRepository
}

func NewTagUsecase(tagRepo repositories.TagRepository, postRepo repositories.PostRepository) TagUsecase {
	return &tagUsecase{tagRepo: tagRepo, postRepo: postRepo}
}

func (s *tagUsecase) CreateTag(name string) (*models.Tag, error) {
	// Check if tag already exists
	if _, err := s.tagRepo.GetTagByName(name); err == nil {
		return nil, errors.New("tag with this name already exists")
	}

	tag := &models.Tag{Name: name}
	if err := s.tagRepo.CreateTag(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *tagUsecase) GetTagByID(id uint) (*models.Tag, error) {
	return s.tagRepo.GetTagByID(id)
}

func (s *tagUsecase) GetTagByName(name string) (*models.Tag, error) {
	return s.tagRepo.GetTagByName(name)
}

func (s *tagUsecase) GetAllTags() ([]models.Tag, error) {
	return s.tagRepo.GetAllTags()
}

func (s *tagUsecase) UpdateTag(id uint, name string) (*models.Tag, error) {
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

func (s *tagUsecase) DeleteTag(id uint) error {
	return s.tagRepo.DeleteTag(id)
}

func (s *tagUsecase) AddTagsToPost(postID uint, tagNames []string) error {
	var tagIDs []uint
	for _, tagName := range tagNames {
		tag, err := s.tagRepo.GetTagByName(tagName)
		if err != nil {
			// Tag does not exist, create it
			newTag := &models.Tag{Name: tagName}
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

func (s *tagUsecase) RemoveTagsFromPost(postID uint, tagNames []string) error {
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

func (s *tagUsecase) GetTagsByPostID(postID uint) ([]models.Tag, error) {
	return s.tagRepo.GetTagsByPostID(postID)
}
