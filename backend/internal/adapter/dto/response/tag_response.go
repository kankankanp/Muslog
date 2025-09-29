package response

import (
	"github.com/kankankanp/Muslog/internal/domain/entity"
)

// タグレスポンス
type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToTagResponse(t *entity.Tag) TagResponse {
	return TagResponse{
		ID:   t.ID,
		Name: t.Name,
	}
}

func ToTagResponses(tags []entity.Tag) []TagResponse {
	res := make([]TagResponse, 0, len(tags))
	for _, t := range tags {
		res = append(res, ToTagResponse(&t))
	}
	return res
}

// 単一取得
type TagDetailResponse struct {
	Message string      `json:"message"`
	Tag     TagResponse `json:"tag"`
}

// 一覧取得
type TagListResponse struct {
	Message string        `json:"message"`
	Tags    []TagResponse `json:"tags"`
}
