package engine

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Publisher interface {
	PublishCalculatedBets(ctx context.Context, bets <-chan rabbitmqmodels.BetCalculated)
}
