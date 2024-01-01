package service

import (
	"context"

	"github.com/ahmtsenlik/UserNotification/internal/pkg/rabbitmq"
	"github.com/ahmtsenlik/UserNotification/internal/userservice/event"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

// NotificationService is the interface that defines the methods of the notification service
type NotificationService interface {
	// SendNotification sends a notification to a user
	SendNotification(ctx context.Context, user *event.UserCreated) error
	// ListNotifications returns all the notifications
	ListNotifications(ctx context.Context) ([]*Notification, error)
}

// notificationService is the default implementation of the NotificationService interface
type notificationService struct {
	notifications []*Notification // A slice of notifications to store the notifications in memory
	broker        broker.Broker   // A broker to publish and subscribe to events
	topic         string          // A topic to publish and subscribe to events
}

// New returns a new instance of the notificationService
func New() (NotificationService, error) {
	// Create a new broker using RabbitMQ
	b := rabbitmq.NewBroker(
		broker.Addrs("amqp://guest:guest@localhost:5672"),
	)

	// Connect to the broker
	if err := b.Connect(); err != nil {
		return nil, err
	}

	// Create a new notification service with an empty slice of notifications, the broker and the topic
	ns := &notificationService{
		notifications: make([]*Notification, 0),
		broker:        b,
		topic:         "user.created",
	}

	// Subscribe to the user created events from the broker with the topic
	if _, err := micro.SubscribeEvent(context.Background(), ns.broker, ns.topic, ns.handleUserCreatedEvent); err != nil {
		return nil, err
	}

	// Return the notification service
	return ns, nil
}

// SendNotification sends a notification to a user
func (s *notificationService) SendNotification(ctx context.Context, user *event.UserCreated) error {
	// Create a new notification with the user ID, name and email
	notification := &Notification{
		UserID:    user.Id,
		UserName:  user.Name,
		UserEmail: user.Email,
		Message:   "Welcome to our platform!",
	}

	// Append the notification to the slice of notifications
	s.notifications = append(s.notifications, notification)

	// Create a new notification sent event with the notification
	e := &event.NotificationSent{
		Notification: notification,
	}

	// Publish the notification sent event to the broker with the topic
	if err := micro.PublishEvent(ctx, s.broker, s.topic, e); err != nil {
		return err
	}

	// Return nil
	return nil
}

// ListNotifications returns all the notifications
func (s *notificationService) ListNotifications(ctx context.Context) ([]*Notification, error) {
	// Return the slice of notifications
	return s.notifications, nil
}

// handleUserCreatedEvent is the function that handles the user created events from the broker
func (s *notificationService) handleUserCreatedEvent(ctx context.Context, user *event.UserCreated) error {
	// Call the notification service to send a notification to the user
	if err := s.SendNotification(ctx, user); err != nil {
		return err
	}

	// Return nil
	return nil
}
