package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/superbet-group/code-cadets-2021/lecture_4/01_event_api/internal/api/controllers/models"
)

// Controller implements handlers for web server requests.
type Controller struct {
	eventUpdateValidator EventUpdateValidator
	eventService         EventService
}

// NewController creates a new instance of Controller
func NewController(eventUpdateValidator EventUpdateValidator, eventService EventService) *Controller {
	return &Controller{
		eventUpdateValidator: eventUpdateValidator,
		eventService:         eventService,
	}
}

// UpdateEvent handlers update event equest.
func (e *Controller) UpdateEvent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var eventUpdateRequestDto models.EventUpdateRequestDto
		err := ctx.ShouldBindWith(&eventUpdateRequestDto, binding.JSON)
		if err != nil {
			ctx.String(http.StatusBadRequest, "update request is not valid.")
			return
		}

		if !e.eventUpdateValidator.EventUpdateIsValid(eventUpdateRequestDto) {
			ctx.String(http.StatusBadRequest, "update request is not valid.")
			return
		}

		err = e.eventService.UpdateEvent(eventUpdateRequestDto.Id, eventUpdateRequestDto.Outcome)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "request could not be processed.")
			return
		}

		ctx.Status(http.StatusOK)
	}
}
