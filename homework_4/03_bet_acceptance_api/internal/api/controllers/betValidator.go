package controllers

import "github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/internal/api/controllers/models"

type BetValidator interface {
	BetIsValid(betDto models.BetDto) bool
}
