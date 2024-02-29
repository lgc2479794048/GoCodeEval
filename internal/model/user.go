package model

// User defines the structure for the user model.
type User struct {
	ID       string
	Email    string
	Username string
	Password string
}

// UserResponse defines the response structure for user-related methods.
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
