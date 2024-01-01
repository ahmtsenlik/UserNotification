package service

// User is the struct that defines a user
type User struct {
	ID    string `json:"id"`    // The user ID
	Name  string `json:"name"`  // The user name
	Email string `json:"email"` // The user email
}
