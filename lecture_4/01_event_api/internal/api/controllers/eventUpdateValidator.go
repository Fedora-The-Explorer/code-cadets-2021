package controllers

import "github.com/superbet-group/code-cadets-2021/lecture_4/01_event_api/internal/api/controllers/models"

// EventUpdateValidator validates event update requests.
type EventUpdateValidator interface {
	EventUpdateIsValid(eventUpdateRequestDto models.EventUpdateRequestDto) bool
}
