package services

// EventService implements event related functions.
type EventService struct {
	eventUpdatePublisher EventUpdatePublisher
}

// NewEventService creates a new instance of EventService.
func NewEventService(eventUpdatePublisher EventUpdatePublisher) *EventService {
	return &EventService{
		eventUpdatePublisher: eventUpdatePublisher,
	}
}

// UpdateEvent sends event update message to the queues.
func (e EventService) UpdateEvent(eventId string, outcome string) error {
	return e.eventUpdatePublisher.Publish(eventId, outcome)
}
