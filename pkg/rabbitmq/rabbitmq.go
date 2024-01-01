package rabbitmq

import (
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-plugins/broker/rabbitmq/v2"
)

// NewBroker returns a new broker using RabbitMQ
func NewBroker(opts ...broker.Option) broker.Broker {
	return rabbitmq.NewBroker(opts...)
}
