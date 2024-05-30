package aggregate

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/events"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/pkg/errors"
)

const (
	ProfileAggregateType es.AggregateType = "profile"
)

type ProfileAggregate struct {
	*es.AggregateBase
	Profile *models.Profile
}

func NewProfileAggregateWithID(id string) *ProfileAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewProfileAggregate()
	aggregate.SetID(id)
	return aggregate
}

func NewProfileAggregate() *ProfileAggregate {
	profileAggregate := &ProfileAggregate{}
	base := es.NewAggregateBase(profileAggregate.When)
	base.SetType(ProfileAggregateType)
	profileAggregate.AggregateBase = base
	return profileAggregate
}

func (a *ProfileAggregate) When(evt es.Event) error {
	switch evt.GetEventType() {

	case events.ProfileCreated:
		return a.onProfileCreated(evt)
	case events.ProfileUpdated:
		return a.onProfileUpdated(evt)

	default:
		return es.ErrInvalidEventType
	}
}

func (a *ProfileAggregate) onProfileCreated(evt es.Event) error {
	var eventData events.ProfileCreatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.Profile = &eventData.Profile
	return nil
}

func (a *ProfileAggregate) onProfileUpdated(evt es.Event) error {
	var eventData events.ProfileUpdatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.Profile = &eventData.Profile
	return nil
}
