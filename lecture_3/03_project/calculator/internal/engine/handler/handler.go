package handler

import (
	"context"
	"log"

	domainmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/domain/models"
	rabbitmqmodels "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/infrastructure/rabbitmq/models"
)

// Handler handles bets received and bets calculated.
type Handler struct {
	betRepository BetRepository
}

// New creates and returns a new Handler.
func New(betRepository BetRepository) *Handler {
	return &Handler{
		betRepository: betRepository,
	}
}

// HandleBetsFromController handles bets from controller.
func (h *Handler) HandleBetsFromController(
	ctx context.Context,
	betsFromController <-chan rabbitmqmodels.BetFromController,
) <-chan rabbitmqmodels.BetCalculated {
	betsCalculated := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(betsCalculated)

		for bet := range betsFromController {
			log.Println("Processing bet received, betId:", bet.Id)

			domainBet := domainmodels.Bet{
				Id:                   bet.Id,
				SelectionId:          bet.SelectionId,
				SelectionCoefficient: bet.SelectionCoefficient,
				Payment:              bet.Payment,
			}

			err := h.betRepository.InsertBet(ctx, domainBet)
			if err != nil {
				log.Println("Failed to insert bet, error: ", err)
				continue
			}
		}
	}()

	return betsCalculated
}

// HandleEventUpdates handles updated bets.
func (h *Handler) HandleEventUpdates(
	ctx context.Context,
	eventUpdates <-chan rabbitmqmodels.BetEventUpdate,
) <-chan rabbitmqmodels.BetCalculated {
	betsCalculated := make(chan rabbitmqmodels.BetCalculated)

	go func() {
		defer close(betsCalculated)

		for eventUpdate := range eventUpdates {
			log.Println("Processing event updates, betId:", eventUpdate.Id)

			// Fetch the domain bet.
			domainBets, exists, err := h.betRepository.GetBetsBySelectionID(ctx, eventUpdate.Id)
			if err != nil {
				log.Println("Failed to fetch a bet which should be updated, error: ", err)
				continue
			}
			if !exists {
				log.Println("A bet which should be updated does not exist, betId: ", eventUpdate.Id)
				continue
			}

			for _, domainBet := range domainBets {
				// Calculate the resulting bet, which should be published.
				betCalculated := rabbitmqmodels.BetCalculated{
					Id: domainBet.Id,
				}

				if eventUpdate.Outcome == "won" {
					betCalculated.Status = "won"
					betCalculated.Payout = domainBet.Payment * domainBet.SelectionCoefficient
				} else if eventUpdate.Outcome == "lost" {
					betCalculated.Status = "lost"
					betCalculated.Payout = 0
				} else {
					log.Println("given event updates do not exist: ", eventUpdate.Id)
					break
				}

				select {
				case betsCalculated <- betCalculated:
				case <-ctx.Done():
					return
				}
			}
		}
	}()

	return betsCalculated
}
