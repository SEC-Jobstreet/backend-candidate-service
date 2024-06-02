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

func (a *ApplicationAggregate) ApplyJob(ctx context.Context, application models.Application) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "ApplicationAggregate.ApplyJob")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", a.GetID()))

	event, err := events.NewJobAppliedEvent(a, application)
	if err != nil {
		return errors.Wrap(err, "NewJobAppliedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}
