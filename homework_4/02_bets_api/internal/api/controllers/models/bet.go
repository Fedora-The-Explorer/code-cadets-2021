package models

// Bet is a storage model representation of a bet.
type Bet struct {
	Id                   string
	CustomerId           string
	Status               string
	SelectionId          string
	SelectionCoefficient float64
	Payment              float64
	Payout               float64
}
