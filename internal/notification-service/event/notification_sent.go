package event

import "github.com/ahmtsenlik/UserNotification/internal/notificationservice/service"

// NotificationSent is the struct that defines the notification sent event
type NotificationSent struct {
	Notification *service.Notification `json:"notification"` // The notification
}
