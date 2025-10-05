package usecase

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
)

type MessageUsecase interface {
	SaveMessage(message *entity.Message) error
	GetMessagesByCommunityID(communityID string) ([]*entity.Message, error)
}

type messageUsecaseImpl struct {
	repo domainRepo.MessageRepository
}

func NewMessageUsecase(repo domainRepo.MessageRepository) MessageUsecase {
	return &messageUsecaseImpl{repo: repo}
}

func (u *messageUsecaseImpl) SaveMessage(message *entity.Message) error {
	return u.repo.Save(message)
}

func (u *messageUsecaseImpl) GetMessagesByCommunityID(communityID string) ([]*entity.Message, error) {
	return u.repo.FindByCommunityID(communityID)
}
