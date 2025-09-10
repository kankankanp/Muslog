package repository

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"gorm.io/gorm"
)

// messageRepositoryImpl implements MessageRepository using GORM.
type messageRepositoryImpl struct {
	DB *gorm.DB
}

// NewMessageRepository creates a new MessageRepository.
func NewMessageRepository(db *gorm.DB) domainRepo.MessageRepository {
	return &messageRepositoryImpl{DB: db}
}

// Save saves a message to the database.
func (r *messageRepositoryImpl) Save(message *entity.Message) error {
	return r.DB.Create(message).Error
}

// FindByCommunityID retrieves messages for a given community ID, ordered by creation time.
func (r *messageRepositoryImpl) FindByCommunityID(communityID string) ([]entity.Message, error) {
	var messages []entity.Message
	err := r.DB.Where("community_id = ?", communityID).
		Order("created_at ASC").
		Find(&messages).Error
	return messages, err
}
