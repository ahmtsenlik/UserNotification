package user

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users []User

func SaveUser(u User) error {
	users = append(users, u)
	return nil
}

func GetUserByID(userID string) (User, error) {
	for _, u := range users {
		if u.ID == userID {
			return u, nil
		}
	}
	return User{}, nil
}

func DeleteUser(userID string) error {
	for i, u := range users {
		if u.ID == userID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return nil
}
