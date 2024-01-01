package main

import (
	"log"
	"net/http"
	"user-service/internal/handlers"
)

func main() {
	r := handlers.SetupRouter()

	log.Println("Server is running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
