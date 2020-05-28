package domain

// Post .
type Post struct {
	UserID    string `json:"userId"`
	Text      string `json:"text"`
	Title     string `json:"title"`
	ID        string `json:"id"`
	CreatedAt int64  `json:"createdAt"`
}
