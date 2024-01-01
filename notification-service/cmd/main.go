package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	notification "notification-service/internal"

	"github.com/streadway/amqp"
)

func main() {
	notificationService := notification.NewService()

	// RabbitMQ connection setup
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	exchangeName := "user_exchange"
	err = ch.ExchangeDeclare(
		exchangeName,        // exchange name
		amqp.ExchangeFanout, // exchange type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	queueName := "user_queue"
	_, err = ch.QueueDeclare(
		queueName, // queue name (should be the same as user service)
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Bind the queue to the exchange
	err = ch.QueueBind(
		queueName,    // queue name
		"",           // routing key
		exchangeName, // exchange name
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to bind the queue to the exchange: %v", err)
	}

	go notificationService.ListenForUserCreationEvents(ch, exchangeName, "")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
