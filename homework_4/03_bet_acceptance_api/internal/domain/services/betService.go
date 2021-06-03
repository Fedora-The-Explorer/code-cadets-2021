package services

import (
	"github.com/nu7hatch/gouuid"
	"log"
)

type BetService struct {
	betPublisher BetPublisher
}

func NewBetService(publisher BetPublisher) *BetService {
	return &BetService{
		betPublisher: publisher,
	}
}

// Publisher gives the bet an id and sends bet message to the queue.
func (e BetService) Publisher(customerId string, selectionId string, selectionCoefficient float64, payment float64) error {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("%s: %s", "failed to create uuid", err)
	}

	return e.betPublisher.Publish(id.String(), customerId, selectionId, selectionCoefficient, payment)
}
