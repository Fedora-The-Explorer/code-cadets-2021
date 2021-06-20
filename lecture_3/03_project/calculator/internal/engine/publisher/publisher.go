package publisher

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Publisher offers methods for publishing into output queues.
type Publisher struct {
	betCalculatedPublisher BetCalculatedPublisher
}

// New creates and returns a new Publisher.
func New(betCalculatedPublisher BetCalculatedPublisher) *Publisher {
	return &Publisher{
		betCalculatedPublisher: betCalculatedPublisher,
	}
}

// PublishCalculatedBets publishes calculated bets into the queue.
func (p *Publisher) PublishCalculatedBets(ctx context.Context, bets <-chan rabbitmqmodels.BetCalculated) {
	p.betCalculatedPublisher.Publish(ctx, bets)
}
