package domain

// User represents our response payload for users
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
