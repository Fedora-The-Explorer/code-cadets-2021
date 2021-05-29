package controllers

// EventService implements event related functions.
type EventService interface {
	UpdateEvent(eventId string, outcome string) error
}
