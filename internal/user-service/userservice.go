package userservice

import (
	"log"

	"github.com/ahmtsenlik/UserNotification/internal/userservice/route"
	"github.com/ahmtsenlik/UserNotification/internal/userservice/service"

	"github.com/micro/go-micro/v2"
)

// Run is the function that runs the user service
func Run() {
	// Create a new service with the name "userservice"
	srv := micro.NewService(
		micro.Name("userservice"),
	)

	// Initialize the service
	srv.Init()

	// Create a new user service instance
	userService, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	// Register the routes for the user service
	r := route.RegisterRoutes(userService)

	// Run the service as a web server
	if err := micro.RegisterHandler(srv.Server(), r); err != nil {
		log.Fatal(err)
	}

	// Run the service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
