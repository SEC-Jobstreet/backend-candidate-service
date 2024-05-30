package postgres_projection

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/events"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

func (o *postgresProjection) onProfileCreate(ctx context.Context, evt es.Event) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "postgresProjection.onProfileCreate")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData events.ProfileCreatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	profile := eventData.Profile
	err := o.postgresRepo.Create(&profile).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *postgresProjection) onProfileUpdate(ctx context.Context, evt es.Event) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "postgresProjection.onProfileUpdate")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()))

	var eventData events.ProfileUpdatedEvent
	if err := evt.GetJsonData(&eventData); err != nil {
		tracing.TraceErr(span, err)
		return errors.Wrap(err, "evt.GetJsonData")
	}

	form := eventData.Profile

	profile := map[string]interface{}{}
	profile["first_name"] = form.FirstName
	profile["last_name"] = form.LastName
	profile["country_phone"] = form.CountryPhone
	profile["phone"] = form.Phone
	profile["address"] = form.Address
	profile["latitude"] = form.Latitude
	profile["longitude"] = form.Longitude
	profile["visa"] = form.Visa
	profile["description"] = form.Description
	profile["current_position"] = form.CurrentPosition
	profile["start_date"] = form.StartDate
	profile["work_whenever"] = form.WorkWhenever
	profile["work_shift"] = form.WorkShift
	profile["share_profile"] = form.ShareProfile

	if form.ResumeLink != "" && form.ResumeName != "" {
		profile["resume_name"] = form.ResumeLink
		profile["resume_link"] = form.ResumeName
	}

	return o.postgresRepo.Model(&models.Profile{Username: form.Username}).Updates(profile).Error
}
