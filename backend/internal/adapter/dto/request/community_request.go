package request

// コミュニティ作成リクエスト
type CreateCommunityRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
