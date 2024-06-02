package aggregate

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/events"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/pkg/errors"
)

const (
	ApplicationAggregateType es.AggregateType = "application"
)

type ApplicationAggregate struct {
	*es.AggregateBase
	Application *models.Application
}

func NewApplicationAggregateWithID(id string) *ApplicationAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewApplicationAggregate()
	aggregate.SetID(id)
	return aggregate
}

func NewApplicationAggregate() *ApplicationAggregate {
	applicationAggregate := &ApplicationAggregate{}
	base := es.NewAggregateBase(applicationAggregate.When)
	base.SetType(ApplicationAggregateType)
	applicationAggregate.AggregateBase = base
	return applicationAggregate
}

func (a *ApplicationAggregate) When(evt es.Event) error {
	switch evt.GetEventType() {

	case events.JobApplied:
		return a.onJobApplied(evt)

	default:
		return es.ErrInvalidEventType
	}
}

func (a *ApplicationAggregate) onJobApplied(evt es.Event) error {
	var eventData events.JobAppliedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.Application = &eventData.Application
	return nil
}
