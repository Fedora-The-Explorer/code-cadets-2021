package rabbitmq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"github.com/superbet-group/code-cadets-2021/lecture_4/01_event_api/internal/infrastructure/rabbitmq/models"
)

const contentTypeTextPlain = "text/plain"

// EventUpdatePublisher handles event update queue publishing.
type EventUpdatePublisher struct {
	exchange  string
	queueName string
	mandatory bool
	immediate bool
	publisher QueuePublisher
}

// NewEventUpdatePublisher create a new instance of EventUpdatePublisher.
func NewEventUpdatePublisher(
	exchange string,
	queueName string,
	mandatory bool,
	immediate bool,
	publisher QueuePublisher,
) *EventUpdatePublisher {
	return &EventUpdatePublisher{
		exchange:  exchange,
		queueName: queueName,
		mandatory: mandatory,
		immediate: immediate,
		publisher: publisher,
	}
}

// Publish publishes an event update message to the queue.
func (p *EventUpdatePublisher) Publish(eventId, outcome string) error {
	eventUpdate := &models.EventUpdateDto{
		Id:      eventId,
		Outcome: outcome,
	}

	eventUpdateJson, err := json.Marshal(eventUpdate)
	if err != nil {
		return err
	}

	err = p.publisher.Publish(
		p.exchange,
		p.queueName,
		p.mandatory,
		p.immediate,
		amqp.Publishing{
			ContentType: contentTypeTextPlain,
			Body:        eventUpdateJson,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent %s", eventUpdateJson)
	return nil
}
