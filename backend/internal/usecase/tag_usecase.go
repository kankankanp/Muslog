package usecase

import (
	"errors"

	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"github.com/kankankanp/Muslog/internal/domain/entity"
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

type tagUsecase struct {
	tagRepo  domainRepo.TagRepository
	postRepo domainRepo.PostRepository
}

func NewTagUsecase(tagRepo domainRepo.TagRepository, postRepo domainRepo.PostRepository) TagUsecase {
	return &tagUsecase{tagRepo: tagRepo, postRepo: postRepo}
}

func (s *tagUsecase) CreateTag(name string) (*entity.Tag, error) {
	if _, err := s.tagRepo.GetTagByName(name); err == nil {
		return nil, errors.New("tag with this name already exists")
	}

	tag := &entity.Tag{Name: name}
	if err := s.tagRepo.CreateTag(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *tagUsecase) GetTagByID(id uint) (*entity.Tag, error) {
	return s.tagRepo.GetTagByID(id)
}

func (s *tagUsecase) GetTagByName(name string) (*entity.Tag, error) {
	return s.tagRepo.GetTagByName(name)
}

func (s *tagUsecase) GetAllTags() ([]entity.Tag, error) {
	return s.tagRepo.GetAllTags()
}

func (s *tagUsecase) UpdateTag(id uint, name string) (*entity.Tag, error) {
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

func (s *tagUsecase) DeleteTag(id uint) error {
	return s.tagRepo.DeleteTag(id)
}

func (s *tagUsecase) AddTagsToPost(postID uint, tagNames []string) error {
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

func (s *tagUsecase) RemoveTagsFromPost(postID uint, tagNames []string) error {
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

func (s *tagUsecase) GetTagsByPostID(postID uint) ([]entity.Tag, error) {
	return s.tagRepo.GetTagsByPostID(postID)
}
