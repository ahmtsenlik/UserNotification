package handlers

import (
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/user", CreateUserHandler).Methods("POST")

	r.HandleFunc("/user/{id}", GetUserHandler).Methods("GET")

	r.HandleFunc("/user/{id}", DeleteUserHandler).Methods("DELETE")

	return r
}
