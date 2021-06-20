package validators

import "github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/internal/api/controllers/models"

const maxCoefficient = 10.0
const minPayment = 2.0
const maxPayment = 100.0

type BetValidator struct{}

func NewBetValidator() *BetValidator {
	return &BetValidator{}
}

// BetIsValid checks if bet is valid
// CustomerId and SelectionId cannot be empty
// And other conditions have to be met
func (e *BetValidator) BetIsValid(betDto models.BetDto) bool {
	if betDto.CustomerId != "" && betDto.SelectionId != "" &&
		betDto.SelectionCoefficient <= maxCoefficient &&
		betDto.Payment >= minPayment &&
		betDto.Payment <= maxPayment {
		return true
	}
	return false
}
