package rabbitmq

import "github.com/streadway/amqp"

// QueuePublisher implements methods for pushing messages to queue.
type QueuePublisher interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}
