package controllers

// EventService implements event related functions.
type ResponseService interface {
	betResponse(eventId string, outcome string) error
}
