package schema

import "time"

// 友達追加リクエスト
type AddFriendRequest struct {
	FriendID int `json:"friend_id"`
}

// 友達一覧のレスポンス用
type FriendResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
