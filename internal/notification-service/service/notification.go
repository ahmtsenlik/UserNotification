package service

// Notification is the struct that defines a notification
type Notification struct {
	UserID    string `json:"user_id"`    // The user ID
	UserName  string `json:"user_name"`  // The user name
	UserEmail string `json:"user_email"` // The user email
	Message   string `json:"message"`    // The notification message
}
