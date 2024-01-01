package handler

import (
	"net/http"

	"github.com/ahmtsenlik/UserNotification/internal/notificationservice/service"
	"github.com/ahmtsenlik/UserNotification/internal/userservice/event"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// SendNotificationRequest is the struct that defines the request body for the send notification handler
type SendNotificationRequest struct {
	UserID    string `json:"user_id" validate:"required"`    // The user ID
	UserName  string `json:"user_name" validate:"required"`  // The user name
	UserEmail string `json:"user_email" validate:"required"` // The user email
}

// SendNotificationResponse is the struct that defines the response body for the send notification handler
type SendNotificationResponse struct {
	Notification *service.Notification `json:"notification"` // The notification
}

// SendNotificationHandler is the handler that sends a notification to a user
func SendNotificationHandler(s service.NotificationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Bind the request body to a SendNotificationRequest struct
		var req SendNotificationRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate the request body
		validate := validator.New()
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create a new user created event with the user ID, name and email
		user := &event.UserCreated{
			Id:    req.UserID,
			Name:  req.UserName,
			Email: req.UserEmail,
		}

		// Call the notification service to send a notification to the user
		if err := s.SendNotification(c.Request.Context(), user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create a new notification with the user ID, name, email and message
		notification := &service.Notification{
			UserID:    user.Id,
			UserName:  user.Name,
			UserEmail: user.Email,
			Message:   "Welcome to our platform!",
		}

		// Create a SendNotificationResponse with the notification
		res := &SendNotificationResponse{
			Notification: notification,
		}

		// Return the response with status code 200
		c.JSON(http.StatusOK, res)
	}
}
