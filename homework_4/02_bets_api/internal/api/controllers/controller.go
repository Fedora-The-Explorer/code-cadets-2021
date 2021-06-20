package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller implements handlers for web server requests.
type Controller struct {
	betResponse BetResponse
}

// NewController creates a new instance of Controller
func NewController(betResponse BetResponse) *Controller {
	return &Controller{
		betResponse: betResponse,
	}
}
func (e *Controller) HandleBetById() gin.HandlerFunc{
	return func(ctx *gin.Context){
		id := ctx.Param("id")
		bet, exists, err := e.betResponse.GetBetById(ctx, id)

		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		if !exists {
			ctx.String(http.StatusNotFound, "failed to get bets with given id")
			return
		}

		ctx.JSON(http.StatusOK, bet)
	}
}

func (e *Controller) HandleBetsByUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		bets, err := e.betResponse.GetBetsByUser(ctx, userId)

		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		if len(bets) == 0 {
			ctx.String(http.StatusNotFound, "failed to get bets with given user id")
			return
		}

		ctx.JSON(http.StatusOK, bets)
	}
}

func (e *Controller) HandleBetsByStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := ctx.Query("status")
		bets, err := e.betResponse.GetBetsByStatus(ctx, status)

		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		if len(bets) == 0 {
			ctx.String(http.StatusNotFound, "failed to get bets with given status")
			return
		}

		ctx.JSON(http.StatusOK, bets)
	}
}
