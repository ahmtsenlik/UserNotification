package notificationservice

import (
	"log"

	"github.com/ahmtsenlik/UserNotification/internal/notificationservice/route"
	"github.com/ahmtsenlik/UserNotification/internal/notificationservice/service"

	"github.com/micro/go-micro/v2"
)

// Run is the function that runs the notification service
func Run() {
	// Create a new service with the name "notificationservice"
	srv := micro.NewService(
		micro.Name("notificationservice"),
	)

	// Initialize the service
	srv.Init()

	// Create a new notification service instance
	notificationService, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	// Register the routes for the notification service
	r := route.RegisterRoutes(notificationService)

	// Run the service as a web server
	if err := micro.RegisterHandler(srv.Server(), r); err != nil {
		log.Fatal(err)
	}

	// Run the service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
