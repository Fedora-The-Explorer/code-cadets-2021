package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pkg/errors"

	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// BetEventUpdateConsumer consumes calculated bets from the desired RabbitMQ queue.
type BetEventUpdateConsumer struct {
	channel Channel
	config  ConsumerConfig
}

// NewBetEventUpdateConsumer creates and returns a new BetEventUpdateConsumer.
func NewBetEventUpdateConsumer(channel Channel, config ConsumerConfig) (*BetEventUpdateConsumer, error) {
	_, err := channel.QueueDeclare(
		config.Queue,
		config.DeclareDurable,
		config.DeclareAutoDelete,
		config.DeclareExclusive,
		config.DeclareNoWait,
		config.DeclareArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet calculated consumer initialization failed")
	}

	return &BetEventUpdateConsumer{
		channel: channel,
		config:  config,
	}, nil
}

// Consume consumes messages until the context is cancelled. An error will be returned if consuming
// is not possible.
func (c *BetEventUpdateConsumer) Consume(ctx context.Context) (<-chan models.BetEventUpdate, error) {
	msgs, err := c.channel.Consume(
		c.config.Queue,
		c.config.ConsumerName,
		c.config.AutoAck,
		c.config.Exclusive,
		c.config.NoLocal,
		c.config.NoWait,
		c.config.Args,
	)
	if err != nil {
		return nil, errors.Wrap(err, "event update consumer failed to consume messages")
	}

	eventUpdates := make(chan models.BetEventUpdate)
	go func() {
		defer close(eventUpdates)
		for msg := range msgs {
			var eventUpdate models.BetEventUpdate
			err := json.Unmarshal(msg.Body, &eventUpdate)
			if err != nil {
				log.Println("Failed to unmarshal bet received message", msg.Body)
			}

			// Once context is cancelled, stop consuming messages.
			select {
			case eventUpdates <- eventUpdate:
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventUpdates, nil
}
