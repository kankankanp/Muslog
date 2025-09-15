package repository

import "github.com/kankankanp/Muslog/internal/domain/entity"

type MessageRepository interface {
	Save(message *entity.Message) error
	FindByCommunityID(communityID string) ([]entity.Message, error)
}
