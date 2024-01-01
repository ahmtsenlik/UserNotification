package handlers

import (
	"encoding/json"
	"net/http"
	"user-service/internal/rabbitmq"
	"user-service/internal/user"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := user.SaveUser(newUser)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	err = rabbitmq.PublishUserEvent(newUser)
	if err != nil {
		http.Error(w, "Error publishing user event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
