package engine

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Consumer interface {
	ConsumeBetsFromController(ctx context.Context) (<-chan rabbitmqmodels.BetFromController, error)
	ConsumeEventUpdates(ctx context.Context) (<-chan rabbitmqmodels.BetEventUpdate, error)
}
