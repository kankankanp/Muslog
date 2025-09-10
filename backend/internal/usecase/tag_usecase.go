package usecase

import (
	"errors"

	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type TagService interface {
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

type tagService struct {
	tagRepo  domainRepo.TagRepository
	postRepo domainRepo.PostRepository
}

func NewTagService(tagRepo domainRepo.TagRepository, postRepo domainRepo.PostRepository) TagService {
	return &tagService{tagRepo: tagRepo, postRepo: postRepo}
}

func (s *tagService) CreateTag(name string) (*entity.Tag, error) {
	if _, err := s.tagRepo.GetTagByName(name); err == nil {
		return nil, errors.New("tag with this name already exists")
	}

	tag := &entity.Tag{Name: name}
	if err := s.tagRepo.CreateTag(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *tagService) GetTagByID(id uint) (*entity.Tag, error) {
	return s.tagRepo.GetTagByID(id)
}

func (s *tagService) GetTagByName(name string) (*entity.Tag, error) {
	return s.tagRepo.GetTagByName(name)
}

func (s *tagService) GetAllTags() ([]entity.Tag, error) {
	return s.tagRepo.GetAllTags()
}

func (s *tagService) UpdateTag(id uint, name string) (*entity.Tag, error) {
	tag, err := s.tagRepo.GetTagByID(id)
	if err != nil {
		return nil, err
	}
	tag.Name = name
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
			newTag := &entity.Tag{Name: tagName}
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
			continue
		}
		tagIDs = append(tagIDs, tag.ID)
	}
	return s.tagRepo.RemoveTagsFromPost(postID, tagIDs)
}

func (s *tagService) GetTagsByPostID(postID uint) ([]entity.Tag, error) {
	return s.tagRepo.GetTagsByPostID(postID)
}
