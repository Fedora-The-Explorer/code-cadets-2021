package controllers

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/api/controllers/models"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/infrastructure/sqlite"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller implements handlers for web server requests.
type Controller struct {
	//eventUpdateValidator EventUpdateValidator
	//eventService         EventService

}

func (e *Controller) GetBetsByUser() gin.HandlerFunc {
	panic("implement me")
}

func (e *Controller) GetBetsByStatus() gin.HandlerFunc {
	panic("implement me")
}

// NewController creates a new instance of Controller
func NewController() *Controller {
	return &Controller{
		//eventUpdateValidator: eventUpdateValidator,
		//eventService:         eventService,
	}
}

// UpdateEvent handlers update event equest.
func (e *Controller) GetBetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		betReceived, err := e.
		// dohhvat
		bet := models.BetResponseDto{
			Id:                   id,
			Status:               "b",
			SelectionId:          "c",
			SelectionCoefficient: 1,
			Payment:              2,
			Payout:               3,
		}
		ctx.JSON(http.StatusOK,bet)
	}
}
