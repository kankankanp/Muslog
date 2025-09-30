package response

import (
	"time"

	"github.com/kankankanp/Muslog/internal/domain/entity"
)

// 投稿レスポンス
type PostResponse struct {
	ID             uint            `json:"id"`
	Title          string          `json:"title"`
	Description    string          `json:"description"`
	UserID         string          `json:"userId"`
	HeaderImageUrl string          `json:"headerImageUrl"`
	Tracks         []TrackResponse `json:"tracks"`
	Tags           []TagResponse   `json:"tags"`
	LikesCount     int             `json:"likesCount"`
	IsLiked        bool            `json:"isLiked"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

func ToPostResponse(p *entity.Post) PostResponse {
	return PostResponse{
		ID:             p.ID,
		Title:          p.Title,
		Description:    p.Description,
		UserID:         p.UserID,
		HeaderImageUrl: p.HeaderImageUrl,
		Tracks:         ToTrackResponses(p.Tracks),
		Tags:           ToTagResponses(p.Tags),
		LikesCount:     p.LikesCount,
		IsLiked:        p.IsLiked,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func ToPostResponses(posts []*entity.Post) []PostResponse {
	res := make([]PostResponse, 0, len(posts))
	for _, p := range posts {
		res = append(res, ToPostResponse(p))
	}
	return res
}

// 一覧レスポンス
type PostListResponse struct {
	Message    string         `json:"message"`
	Posts      []PostResponse `json:"posts"`
	TotalCount int64          `json:"totalCount,omitempty"`
	Page       int            `json:"page,omitempty"`
	PerPage    int            `json:"perPage,omitempty"`
}

// 単体レスポンス
type PostDetailResponse struct {
	Message string       `json:"message"`
	Post    PostResponse `json:"post"`
}
