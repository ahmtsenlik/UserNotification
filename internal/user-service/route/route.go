package route

import (
	"github.com/ahmtsenlik/UserNotification/internal/userservice/handler"
	"github.com/ahmtsenlik/UserNotification/internal/userservice/service"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for the user service
func RegisterRoutes(s service.UserService) *gin.Engine {
	// Create a new Gin engine
	r := gin.Default()

	// Create a new user group
	userGroup := r.Group("/user")

	// Register the create user handler with the POST method
	userGroup.POST("/", handler.CreateUserHandler(s))

	// Register the list users handler with the GET method
	userGroup.GET("/", handler.ListUsersHandler(s))

	// Return the Gin engine
	return r
}
