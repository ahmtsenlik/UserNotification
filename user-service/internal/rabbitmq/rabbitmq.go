package rabbitmq

import (
	"encoding/json"
	"fmt"
	"user-service/internal/user"

	"github.com/streadway/amqp"
)

const (
	amqpURL      = "amqp://guest:guest@localhost:5672/"
	exchangeName = "user_exchange"
	queueName    = "user_queue"
)

func PublishUserEvent(u user.User) error {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare an exchange: %v", err)
	}

	// User modelini JSON formatına çevir
	jsonBody, err := json.Marshal(u)
	if err != nil {
		return fmt.Errorf("failed to marshal user to JSON: %v", err)
	}

	err = ch.Publish(
		exchangeName, // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		})
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	return nil
}
