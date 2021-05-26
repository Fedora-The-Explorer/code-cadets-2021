package models

// BetEventUpdate represents a DTO for updated bets.
type BetEventUpdate struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}
