package consumer

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type BetEventUpdateConsumer interface {
	Consume(ctx context.Context) (<-chan rabbitmqmodels.BetEventUpdate, error)
}