package handler

import (
	"net/http"

	"github.com/ahmtsenlik/UserNotification/internal/notificationservice/service"
	"github.com/gin-gonic/gin"
)

// ListNotificationsResponse is the struct that defines the response body for the list notifications handler
type ListNotificationsResponse struct {
	Notifications []*service.Notification `json:"notifications"` // The notifications
}

// ListNotificationsHandler is the handler that returns all the notifications
func ListNotificationsHandler(s service.NotificationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call the notification service to list all the notifications
		notifications, err := s.ListNotifications(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create a ListNotificationsResponse with the notifications
		res := &ListNotificationsResponse{
			Notifications: notifications,
		}

		// Return the response with status code 200
		c.JSON(http.StatusOK, res)
	}
}
