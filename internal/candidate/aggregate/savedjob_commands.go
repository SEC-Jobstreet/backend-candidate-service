package aggregate

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/events"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

func (a *SavedJobAggregate) SaveJob(ctx context.Context, savedJob models.SavedJob) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "SavedJobAggregate.SaveJob")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", a.GetID()))

	event, err := events.NewJobSavedEvent(a, savedJob)
	if err != nil {
		return errors.Wrap(err, "NewJobSavedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}

func (a *SavedJobAggregate) UnsaveJob(ctx context.Context, savedJob models.SavedJob) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "SavedJobAggregate.UnsaveJob")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", a.GetID()))

	event, err := events.NewJobUnsavedEvent(a, savedJob)
	if err != nil {
		return errors.Wrap(err, "NewJobUnsavedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
