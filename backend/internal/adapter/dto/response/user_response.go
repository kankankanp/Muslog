package response

import "github.com/kankankanp/Muslog/internal/domain/entity"

// ユーザー情報レスポンス
type UserResponse struct {
	ID              string                `json:"id"`
	Name            string                `json:"name"`
	Email           string                `json:"email"`
	ProfileImageUrl string                `json:"profileImageUrl"`
	Setting         *UserSettingResponse  `json:"setting,omitempty"`
}

func ToUserResponse(u *entity.User) UserResponse {
	userResp := UserResponse{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		ProfileImageUrl: u.ProfileImageUrl,
	}
	if u.Setting != nil {
		userResp.Setting = ToUserSettingResponse(u.Setting)
	}
	return userResp
}

// ログイン・登録成功レスポンス
type AuthResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

// ユーザー一覧レスポンス
type UserListResponse struct {
	Message string         `json:"message"`
	Users   []UserResponse `json:"users"`
}

// 単一ユーザー取得レスポンス
type UserDetailResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

// 投稿一覧レスポンス
type UserPostsResponse struct {
	Message string         `json:"message"`
	Posts   []PostResponse `json:"posts"`
}
