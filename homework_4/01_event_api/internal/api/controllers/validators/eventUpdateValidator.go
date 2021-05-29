package validators

import "github.com/superbet-group/code-cadets-2021/lecture_4/01_event_api/internal/api/controllers/models"

const lostOutcome = "lost"
const wonOutcome = "won"

// EventUpdateValidator validates event update requests.
type EventUpdateValidator struct{}

// NewEventUpdateValidator creates a new instance of EventUpdateValidator.
func NewEventUpdateValidator() *EventUpdateValidator {
	return &EventUpdateValidator{}
}

// EventUpdateIsValid checks if event update is valid.
// Id is not empty
// Outcome is `lost`or `won`
func (e *EventUpdateValidator) EventUpdateIsValid(eventUpdateRequestDto models.EventUpdateRequestDto) bool {
	if eventUpdateRequestDto.Id != "" && (eventUpdateRequestDto.Outcome == lostOutcome || eventUpdateRequestDto.Outcome == wonOutcome) {
		return true
	}

	return false
}
