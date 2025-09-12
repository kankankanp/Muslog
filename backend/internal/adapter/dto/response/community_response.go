package response

import (
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

// コミュニティレスポンス
type CommunityResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatorID   string    `json:"creatorId"`
	CreatedAt   time.Time `json:"createdAt"`
}

func ToCommunityResponse(c *entity.Community) CommunityResponse {
	return CommunityResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatorID:   c.CreatorID,
		CreatedAt:   c.CreatedAt,
	}
}

func ToCommunityResponses(communities []entity.Community) []CommunityResponse {
	res := make([]CommunityResponse, 0, len(communities))
	for _, c := range communities {
		res = append(res, ToCommunityResponse(&c))
	}
	return res
}

// 作成レスポンス
type CreateCommunityResponse struct {
	Message   string            `json:"message"`
	Community CommunityResponse `json:"community"`
}

// 一覧レスポンス
type CommunityListResponse struct {
	Message     string              `json:"message"`
	Communities []CommunityResponse `json:"communities"`
	TotalCount  int64               `json:"totalCount,omitempty"`
	Page        int                 `json:"page,omitempty"`
	PerPage     int                 `json:"perPage,omitempty"`
}
