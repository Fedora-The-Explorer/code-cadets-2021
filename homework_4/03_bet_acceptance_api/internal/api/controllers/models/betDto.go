package models

// EventUpdateRequestDto Update request dto model.
type BetDto struct {
	CustomerId           string  `json:"customerId"`
	SelectionId          string  `json:"selectionId"`
	SelectionCoefficient float64 `json:"selectionCoefficient"`
	Payment              float64 `json:"payment"`
}
