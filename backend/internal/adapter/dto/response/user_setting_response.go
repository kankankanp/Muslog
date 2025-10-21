package response

import (
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

type UserSettingResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userId"`
	EditorType string    `json:"editorType"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UserSettingDetailResponse struct {
	Message string               `json:"message"`
	Setting *UserSettingResponse `json:"setting"`
}

func ToUserSettingResponse(setting *entity.UserSetting) *UserSettingResponse {
	if setting == nil {
		return nil
	}
	return &UserSettingResponse{
		ID:         setting.ID,
		UserID:     setting.UserID,
		EditorType: setting.EditorType,
		CreatedAt:  setting.CreatedAt,
		UpdatedAt:  setting.UpdatedAt,
	}
}