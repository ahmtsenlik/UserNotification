package notification

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ListenForUserCreationEvents(ch *amqp.Channel, exchangeName, queueName string) {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for msg := range msgs {
		var userCreatedEvent UserCreatedEvent
		err := json.Unmarshal(msg.Body, &userCreatedEvent)
		log.Printf("Received message body: %s", string(msg.Body))

		if err != nil {
			log.Printf("Failed to unmarshal user created event: %v", err)
			continue
		}

		fmt.Printf("Notification sent for user created: %s\n", userCreatedEvent.ID)
	}
}

type UserCreatedEvent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
