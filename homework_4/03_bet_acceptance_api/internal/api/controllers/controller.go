package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/internal/api/controllers/models"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betValidator BetValidator
	betService   BetService
}

// NewController creates a new instance of Controller
func NewController(betValidator BetValidator, betService BetService) *Controller {
	return &Controller{
		betValidator: betValidator,
		betService:   betService,
	}
}

func (e *Controller) HandleBet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var betDto models.BetDto
		err := ctx.ShouldBindWith(&betDto, binding.JSON)
		if err != nil {
			ctx.String(http.StatusBadRequest, "invalid bet")
			return
		}

		if !e.betValidator.BetIsValid(betDto) {
			ctx.String(http.StatusBadRequest, "invalid request")
			return
		}

		err = e.betService.Publisher(betDto.CustomerId, betDto.SelectionId, betDto.SelectionCoefficient, betDto.Payment)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}
		ctx.Status(http.StatusOK)
	}
}