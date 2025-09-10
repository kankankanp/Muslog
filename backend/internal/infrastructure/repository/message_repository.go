package repository

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	"gorm.io/gorm"
)

// MessageRepository defines the interface for message data operations.
type MessageRepository interface {
	Save(message *entity.Message) error
	FindByCommunityID(communityID string) ([]entity.Message, error)
}

// messageRepository implements MessageRepository using GORM.
type messageRepository struct {
	DB *gorm.DB
}

// NewMessageRepository creates a new MessageRepository.
func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{DB: db}
}

// Save saves a message to the database.
func (r *messageRepository) Save(message *entity.Message) error {
	return r.DB.Create(message).Error
}

// FindByCommunityID retrieves messages for a given community ID, ordered by creation time.
func (r *messageRepository) FindByCommunityID(communityID string) ([]entity.Message, error) {
	var messages []entity.Message
	err := r.DB.Where("community_id = ?", communityID).Order("created_at ASC").Find(&messages).Error
	return messages, err
}
