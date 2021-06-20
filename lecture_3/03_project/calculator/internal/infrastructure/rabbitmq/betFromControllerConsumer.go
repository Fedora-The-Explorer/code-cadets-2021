package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pkg/errors"

	"github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// BetFromControllerConsumer consumes received bets from the desired RabbitMQ queue.
type BetFromControllerConsumer struct {
	channel Channel
	config  ConsumerConfig
}

// NewBetFromControllerConsumer creates and returns a new BetFromControllerConsumer.
func NewBetFromControllerConsumer(channel Channel, config ConsumerConfig) (*BetFromControllerConsumer, error) {
	_, err := channel.QueueDeclare(
		config.Queue,
		config.DeclareDurable,
		config.DeclareAutoDelete,
		config.DeclareExclusive,
		config.DeclareNoWait,
		config.DeclareArgs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "bet received consumer initialization failed")
	}

	return &BetFromControllerConsumer{
		channel: channel,
		config:  config,
	}, nil
}

// Consume consumes messages until the context is cancelled. An error will be returned if consuming
// is not possible.
func (c *BetFromControllerConsumer) Consume(ctx context.Context) (<-chan models.BetFromController, error) {
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
		return nil, errors.Wrap(err, "bet received consumer failed to consume messages")
	}

	bets := make(chan models.BetFromController)
	go func() {
		defer close(bets)
		for msg := range msgs {
			var betReceived models.BetFromController
			err := json.Unmarshal(msg.Body, &betReceived)
			if err != nil {
				log.Println("Failed to unmarshal bet received message", msg.Body)
			}

			// Once context is cancelled, stop consuming messages.
			select {
			case bets <- betReceived:
			case <-ctx.Done():
				return
			}
		}
	}()

	return bets, nil
}
