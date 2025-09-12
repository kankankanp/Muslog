package response

// 画像アップロード成功レスポンス
type ImageUploadResponse struct {
	Message  string `json:"message"`
	ImageURL string `json:"imageUrl"`
}
