package events

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
)

const (
	ProfileCreated = "PROFILE_CREATED"
	ProfileUpdated = "PROFILE_UPDATED"
)

type ProfileCreatedEvent struct {
	Profile models.Profile `json:"profile" bson:"profile"`
}

func NewProfileCreatedEvent(aggregate es.Aggregate, profile models.Profile) (es.Event, error) {
	eventData := ProfileCreatedEvent{
		Profile: profile,
	}

	event := es.NewBaseEvent(aggregate, ProfileCreated)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}

	return event, nil
}

type ProfileUpdatedEvent struct {
	Profile models.Profile `json:"profile" bson:"profile"`
}

func NewProfileUpdatedEvent(aggregate es.Aggregate, profile models.Profile) (es.Event, error) {
	eventData := ProfileUpdatedEvent{
		Profile: profile,
	}
	event := es.NewBaseEvent(aggregate, ProfileUpdated)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
