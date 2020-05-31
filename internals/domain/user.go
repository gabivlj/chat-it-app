package domain

// User .
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	ImageURL string `json:"imageURL"`
}
