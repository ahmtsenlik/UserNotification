package event

// UserCreated is the struct that defines the user created event
type UserCreated struct {
	Id    string `json:"id"`    // The user ID
	Name  string `json:"name"`  // The user name
	Email string `json:"email"` // The user email
}
