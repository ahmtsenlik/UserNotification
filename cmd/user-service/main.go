package main

import (
	"log"

	"github.com/ahmtsenlik/UserNotification/internal/userservice"
)

func main() {
	// Create a new user service
	service, err := userservice.New()
	if err != nil {
		log.Fatal(err)
	}

	// Run the user service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
