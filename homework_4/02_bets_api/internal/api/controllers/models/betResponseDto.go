package models

// EventUpdateRequestDto Update request dto model.
type BetResponseDto struct {
	Id                   string  `json:"id"`
	Status               string  `json:"status"`
	SelectionId          string  `json:"selection_id"`
	SelectionCoefficient float64 `json:"selection_coefficient"`
	Payment              float64 `json:"payment"`
	Payout               float64 `json:"payout"`
}
