package usecase

import (
	"context"
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
	domainRepo "github.com/kankankanp/Muslog/internal/domain/repository"
)

type UserSettingUsecase interface {
	GetUserSetting(ctx context.Context, userID string) (*entity.UserSetting, error)
	UpdateUserSetting(ctx context.Context, userID string, editorType string) (*entity.UserSetting, error)
}

type userSettingUsecaseImpl struct {
	userSettingRepo domainRepo.UserSettingRepository
}

func NewUserSettingUsecase(userSettingRepo domainRepo.UserSettingRepository) UserSettingUsecase {
	return &userSettingUsecaseImpl{
		userSettingRepo: userSettingRepo,
	}
}

func (u *userSettingUsecaseImpl) GetUserSetting(ctx context.Context, userID string) (*entity.UserSetting, error) {
	setting, err := u.userSettingRepo.FindByUserID(ctx, userID)
	if err != nil {
		// 設定が存在しない場合はデフォルト設定を作成
		defaultSetting := &entity.UserSetting{
			UserID:     userID,
			EditorType: "markdown",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		return u.userSettingRepo.Create(ctx, defaultSetting)
	}
	return setting, nil
}

func (u *userSettingUsecaseImpl) UpdateUserSetting(ctx context.Context, userID string, editorType string) (*entity.UserSetting, error) {
	setting, err := u.userSettingRepo.FindByUserID(ctx, userID)
	if err != nil {
		// 設定が存在しない場合は新規作成
		newSetting := &entity.UserSetting{
			UserID:     userID,
			EditorType: editorType,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		return u.userSettingRepo.Create(ctx, newSetting)
	}

	// 既存設定を更新
	setting.EditorType = editorType
	setting.UpdatedAt = time.Now()

	err = u.userSettingRepo.Update(ctx, setting)
	if err != nil {
		return nil, err
	}

	return setting, nil
}
