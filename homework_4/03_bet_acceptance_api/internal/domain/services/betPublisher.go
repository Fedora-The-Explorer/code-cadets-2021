package services

// EventUpdatePublisher handles event update queue publishing.
type BetPublisher interface {
	Publish(id, customerId, selectionId string, selectionCoefficient, payment float64) error
}
