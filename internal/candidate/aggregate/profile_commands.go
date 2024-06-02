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

func (a *ProfileAggregate) CreateProfile(ctx context.Context, profile models.Profile) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "ProfileAggregate.CreateProfile")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", a.GetID()))

	event, err := events.NewProfileCreatedEvent(a, profile)
	if err != nil {
		return errors.Wrap(err, "NewProfileCreatedEvent")
	}

	if err := event.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(event)
}

func (a *ProfileAggregate) UpdateProfile(ctx context.Context, profile models.Profile) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "ProfileAggregate.UpdateProfile")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", a.GetID()))

	profileUpdatedEvent, err := events.NewProfileUpdatedEvent(a, profile)
	if err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "NewProfileUpdatedEvent")
	}

	if err := profileUpdatedEvent.SetMetadata(tracing.ExtractTextMapCarrier(span.Context())); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "SetMetadata")
	}

	return a.Apply(profileUpdatedEvent)
}
