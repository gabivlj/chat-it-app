package domain

// Message .
type Message struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	Text      string `json:"text"`
	UserID    string `json:"userId"`
	PostID    string `json:"postId"`

	Post *Post `json:"post"`
	User *User `json:"user"`
}

// MessageCount .
type MessageCount struct {
	Total  uint64 `json:"total"`
	PostID string `json:"postId"`
}
