package services

// EventUpdatePublisher handles event update queue publishing.
type EventUpdatePublisher interface {
	Publish(eventId, outcome string) error
}
