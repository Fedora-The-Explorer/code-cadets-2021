package models

// BetEventUpdate represents a DTO for received bets.
type BetEventUpdate struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}
