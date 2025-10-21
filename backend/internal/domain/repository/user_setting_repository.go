package repository

import (
	"context"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type UserSettingRepository interface {
	FindByUserID(ctx context.Context, userID string) (*entity.UserSetting, error)
	Create(ctx context.Context, setting *entity.UserSetting) (*entity.UserSetting, error)
	Update(ctx context.Context, setting *entity.UserSetting) error
}