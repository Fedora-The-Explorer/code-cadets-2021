package models

// Bet represents a domain model representation of a bet.
type Bet struct {
	Id                   string
	SelectionId          string
	SelectionCoefficient float64
	Payment              float64
}
