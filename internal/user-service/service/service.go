package service

import (
	"context"

	"github.com/ahmtsenlik/UserNotification/internal/pkg/rabbitmq"
	"github.com/ahmtsenlik/UserNotification/internal/userservice/event"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

// UserService is the interface that defines the methods of the user service
type UserService interface {
	// CreateUser creates a new user and publishes a user created event
	CreateUser(ctx context.Context, name string, email string) (*User, error)
	// ListUsers returns all the users
	ListUsers(ctx context.Context) ([]*User, error)
}

// userService is the default implementation of the UserService interface
type userService struct {
	users  []*User       // A slice of users to store the users in memory
	broker broker.Broker // A broker to publish and subscribe to events
	topic  string        // A topic to publish and subscribe to events
}

// New returns a new instance of the userService
func New() (UserService, error) {
	// Create a new broker using RabbitMQ
	b := rabbitmq.NewBroker(
		broker.Addrs("amqp://guest:guest@localhost:5672"),
	)

	// Connect to the broker
	if err := b.Connect(); err != nil {
		return nil, err
	}

	// Create a new user service with an empty slice of users, the broker and the topic
	return &userService{
		users:  make([]*User, 0),
		broker: b,
		topic:  "user.created",
	}, nil
}

// CreateUser creates a new user and publishes a user created event
func (s *userService) CreateUser(ctx context.Context, name string, email string) (*User, error) {
	// Generate a random UUID for the user ID
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// Create a new user with the given name, email and the generated ID
	user := &User{
		ID:    id.String(),
		Name:  name,
		Email: email,
	}

	// Append the user to the slice of users
	s.users = append(s.users, user)

	// Create a new user created event with the user ID, name and email
	e := &event.UserCreated{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	// Publish the user created event to the broker with the topic
	if err := micro.PublishEvent(ctx, s.broker, s.topic, e); err != nil {
		return nil, err
	}

	// Return the user
	return user, nil
}

// ListUsers returns all the users
func (s *userService) ListUsers(ctx context.Context) ([]*User, error) {
	// Return the slice of users
	return s.users, nil
}
