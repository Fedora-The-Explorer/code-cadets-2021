package models

// Bet represents a DTO for bets.
type BetFromController struct {
	Id                   string  `json:"id"`
	CustomerId           string  `json:"customerId"`
	Status               string  `json:"status"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
	Payout               float64 `json:"payout,omitempty"`
}
