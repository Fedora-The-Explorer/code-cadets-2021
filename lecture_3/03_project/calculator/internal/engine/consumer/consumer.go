package consumer

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Consumer offers methods for consuming from input queues.
type Consumer struct {
	betFromControllerConsumer   BetFromController
	eventUpdateConsumer BetEventUpdateConsumer
}

// New creates and returns a new Consumer.
func New(betFromController BetFromController, eventUpdateConsumer BetEventUpdateConsumer) *Consumer {
	return &Consumer{
		betFromControllerConsumer:   betFromController,
		eventUpdateConsumer: eventUpdateConsumer,
	}
}

// ConsumeBets consumes bets queue.
func (c *Consumer) ConsumeBetsFromController(ctx context.Context) (<-chan rabbitmqmodels.BetFromController, error) {
	return c.betFromControllerConsumer.Consume(ctx)
}

// ConsumeEventUpdates consumes event updates queue.
func (c *Consumer) ConsumeEventUpdates(ctx context.Context) (<-chan rabbitmqmodels.BetEventUpdate, error) {
	return c.eventUpdateConsumer.Consume(ctx)
}