package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
	"github.com/kankankanp/Muslog/internal/infrastructure/model"
	"gorm.io/gorm"
)

type userSettingRepositoryImpl struct {
	db *gorm.DB
}

func NewUserSettingRepository(db *gorm.DB) domainRepo.UserSettingRepository {
	return &userSettingRepositoryImpl{db: db}
}

func (r *userSettingRepositoryImpl) FindByUserID(ctx context.Context, userID string) (*entity.UserSetting, error) {
	var settingModel model.UserSettingModel
	if err := r.db.Where("user_id = ?", userID).First(&settingModel).Error; err != nil {
		return nil, err
	}
	return toUserSettingEntity(&settingModel), nil
}

func (r *userSettingRepositoryImpl) Create(ctx context.Context, setting *entity.UserSetting) (*entity.UserSetting, error) {
	settingModel := toUserSettingModel(setting)
	settingModel.ID = uuid.New().String()
	if err := r.db.Create(settingModel).Error; err != nil {
		return nil, err
	}
	return toUserSettingEntity(settingModel), nil
}

func (r *userSettingRepositoryImpl) Update(ctx context.Context, setting *entity.UserSetting) error {
	settingModel := toUserSettingModel(setting)
	return r.db.Save(settingModel).Error
}

func toUserSettingEntity(model *model.UserSettingModel) *entity.UserSetting {
	return &entity.UserSetting{
		ID:         model.ID,
		UserID:     model.UserID,
		EditorType: model.EditorType,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	}
}

func toUserSettingModel(entity *entity.UserSetting) *model.UserSettingModel {
	return &model.UserSettingModel{
		ID:         entity.ID,
		UserID:     entity.UserID,
		EditorType: entity.EditorType,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}
