package controllers

type BetService interface {
	Publisher(customerId string, selectionId string, selectionCoefficient float64, payment float64) error
}
