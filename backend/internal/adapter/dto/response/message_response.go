package response

import (
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

// メッセージレスポンス
type MessageResponse struct {
	ID          string    `json:"id"`
	CommunityID string    `json:"communityId"`
	SenderID    string    `json:"senderId"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}

func ToMessageResponse(m *entity.Message) MessageResponse {
	return MessageResponse{
		ID:          m.ID,
		CommunityID: m.CommunityID,
		SenderID:    m.SenderID,
		Content:     m.Content,
		CreatedAt:   m.CreatedAt,
	}
}

func ToMessageResponses(messages []entity.Message) []MessageResponse {
	res := make([]MessageResponse, 0, len(messages))
	for _, m := range messages {
		res = append(res, ToMessageResponse(&m))
	}
	return res
}

// 一覧レスポンス
type MessageListResponse struct {
	Message  string            `json:"message"`
	Messages []MessageResponse `json:"messages"`
}
