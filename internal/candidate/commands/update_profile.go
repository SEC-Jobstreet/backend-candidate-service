package commands

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/aggregate"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type UpdateProfileCommandHandler interface {
	Handle(ctx context.Context, command *UpdateProfileCommand) error
}

type updateProfileHandler struct {
	cfg *utils.Config
	es  es.AggregateStore
}

func NewUpdateProfileHandler(cfg *utils.Config, es es.AggregateStore) *updateProfileHandler {
	return &updateProfileHandler{cfg: cfg, es: es}
}

func (c *updateProfileHandler) Handle(ctx context.Context, command *UpdateProfileCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateProfileHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	profile, err := aggregate.LoadProfileAggregate(ctx, c.es, command.AggregateID)
	if err != nil {
		return err
	}

	if err := profile.UpdateProfile(ctx, command.Profile); err != nil {
		return err
	}

	span.LogFields(log.String("profile", profile.String()))
	return c.es.Save(ctx, profile)
}
