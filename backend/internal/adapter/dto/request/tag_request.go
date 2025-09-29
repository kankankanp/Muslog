package request

// Create/Update 用
type TagRequest struct {
	Name string `json:"name"`
}

// Post にタグを追加/削除するとき
type TagNamesRequest struct {
	TagNames []string `json:"tag_names"`
}
