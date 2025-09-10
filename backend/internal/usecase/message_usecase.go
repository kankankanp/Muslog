package usecase

import (
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

// MessageUsecase defines the interface for message-related business logic.
type MessageUsecase interface {
	SaveMessage(message *entity.Message) error
	GetMessagesByCommunityID(communityID string) ([]entity.Message, error)
}

// messageUsecase implements MessageUsecase.
type messageUsecase struct {
	repo domainRepo.MessageRepository
}

// NewMessageUsecase creates a new MessageUsecase.
func NewMessageUsecase(repo domainRepo.MessageRepository) MessageUsecase {
	return &messageUsecase{repo: repo}
}

// SaveMessage saves a message using the repository.
func (uc *messageUsecase) SaveMessage(message *entity.Message) error {
	return uc.repo.Save(message)
}

// GetMessagesByCommunityID retrieves messages for a given community ID using the repository.
func (uc *messageUsecase) GetMessagesByCommunityID(communityID string) ([]entity.Message, error) {
	return uc.repo.FindByCommunityID(communityID)
}
