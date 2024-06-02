package aggregate

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/events"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/pkg/errors"
)

const (
	SavedJobAggregateType es.AggregateType = "savedjob"
)

type SavedJobAggregate struct {
	*es.AggregateBase
	SavedJob *models.SavedJob
}

func NewSavedJobAggregateWithID(id string) *SavedJobAggregate {
	if id == "" {
		return nil
	}

	aggregate := NewSavedJobAggregate()
	aggregate.SetID(id)
	return aggregate
}

func NewSavedJobAggregate() *SavedJobAggregate {
	applicationAggregate := &SavedJobAggregate{}
	base := es.NewAggregateBase(applicationAggregate.When)
	base.SetType(SavedJobAggregateType)
	applicationAggregate.AggregateBase = base
	return applicationAggregate
}

func (a *SavedJobAggregate) When(evt es.Event) error {
	switch evt.GetEventType() {

	case events.JobSaved:
		return a.onJobSaved(evt)
	case events.JobUnsaved:
		return a.onJobUnsaved(evt)

	default:
		return es.ErrInvalidEventType
	}
}

func (a *SavedJobAggregate) onJobSaved(evt es.Event) error {
	var eventData events.JobSavedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.SavedJob = &eventData.SavedJob
	return nil
}

func (a *SavedJobAggregate) onJobUnsaved(evt es.Event) error {
	var eventData events.JobUnsavedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		return errors.Wrap(err, "GetJsonData")
	}

	a.SavedJob = &eventData.SavedJob
	return nil
}
