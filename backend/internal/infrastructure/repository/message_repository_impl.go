package repository

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"github.com/kankankanp/Muslog/internal/infrastructure/mapper"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"gorm.io/gorm"
)

type messageRepositoryImpl struct {
	DB *gorm.DB
}

func NewMessageRepository(db *gorm.DB) domainRepo.MessageRepository {
	return &messageRepositoryImpl{DB: db}
}

func (r *messageRepositoryImpl) Save(message *entity.Message) error {
	m := mapper.FromMessageEntity(message)
	return r.DB.Create(m).Error
}

func (r *messageRepositoryImpl) FindByCommunityID(communityID string) ([]*entity.Message, error) {
	var models []model.MessageModel
	if err := r.DB.Where("community_id = ?", communityID).
		Order("created_at ASC").
		Find(&models).Error; err != nil {
		return nil, err
	}

	messages := make([]*entity.Message, 0, len(models))
	for _, m := range models {
		messages = append(messages, mapper.ToMessageEntity(&m))
	}
	return messages, nil
}
