package user

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(u User) error {
	return SaveUser(u)
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(userID string) (User, error) {
	return GetUserByID(userID)
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(userID string) error {
	return DeleteUser(userID)
}
