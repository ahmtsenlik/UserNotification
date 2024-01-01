package main

import (
	"log"

	notificationservice "github.com/ahmtsenlik/UserNotification/internal/notification-service"
)

func main() {
	// Create a new notification service
	service, err := notificationservice.New()
	if err != nil {
		log.Fatal(err)
	}

	// Run the notification service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
