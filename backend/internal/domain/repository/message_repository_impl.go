package repository

import "github.com/kankankanp/Muslog/internal/domain/entity"

// MessageRepository defines the contract for message data operations.
type MessageRepository interface {
	Save(message *entity.Message) error
	FindByCommunityID(communityID string) ([]entity.Message, error)
}