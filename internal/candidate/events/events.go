package events

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
)

const (
	ProfileCreated = "PROFILE_CREATED"
	ProfileUpdated = "PROFILE_UPDATED"
	JobApplied     = "JOB_APPLIED"
	JobSaved       = "JOB_SAVED"
	JobUnsaved     = "JOB_UNSAVED"
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

type JobAppliedEvent struct {
	Application models.Application `json:"application" bson:"application"`
}

func NewJobAppliedEvent(aggregate es.Aggregate, application models.Application) (es.Event, error) {
	eventData := JobAppliedEvent{
		Application: application,
	}
	event := es.NewBaseEvent(aggregate, JobApplied)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

type JobSavedEvent struct {
	SavedJob models.SavedJob `json:"savedjob" bson:"savedjob"`
}

func NewJobSavedEvent(aggregate es.Aggregate, savedjob models.SavedJob) (es.Event, error) {
	eventData := JobSavedEvent{
		SavedJob: savedjob,
	}
	event := es.NewBaseEvent(aggregate, JobSaved)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}

type JobUnsavedEvent struct {
	SavedJob models.SavedJob `json:"savedjob" bson:"savedjob"`
}

func NewJobUnsavedEvent(aggregate es.Aggregate, unsavedjob models.SavedJob) (es.Event, error) {
	eventData := JobSavedEvent{
		SavedJob: unsavedjob,
	}
	event := es.NewBaseEvent(aggregate, JobUnsaved)
	if err := event.SetJsonData(&eventData); err != nil {
		return es.Event{}, err
	}
	return event, nil
}
