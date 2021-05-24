package engine

import (
	"context"

	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

type Handler interface {
	HandleBetsFromController(ctx context.Context, betsReceived <-chan rabbitmqmodels.BetFromController) <-chan rabbitmqmodels.BetCalculated
	HandleEventUpdates(ctx context.Context, betsCalculated <-chan rabbitmqmodels.BetEventUpdate) <-chan rabbitmqmodels.BetCalculated
}
