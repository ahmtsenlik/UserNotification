package route

import (
	"github.com/ahmtsenlik/UserNotification/internal/notificationservice/handler"
	"github.com/ahmtsenlik/UserNotification/internal/notificationservice/service"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for the notification service
func RegisterRoutes(s service.NotificationService) *gin.Engine {
	// Create a new Gin engine
	r := gin.Default()

	// Create a new notification group
	notificationGroup := r.Group("/notification")

	// Register the send notification handler with the POST method
	notificationGroup.POST("/", handler.SendNotificationHandler(s))

	// Register the list notifications handler with the GET method
	notificationGroup.GET("/", handler.ListNotificationsHandler(s))

	// Return the Gin engine
	return r
}
