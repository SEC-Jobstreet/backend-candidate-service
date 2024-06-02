package commands

import (
	"context"
	"errors"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/aggregate"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type ApplyJobCommandHandler interface {
	Handle(ctx context.Context, command *ApplyJobCommand) error
}

type applyJobHandler struct {
	cfg *utils.Config
	es  es.AggregateStore
}

func NewApplyJobHandler(cfg *utils.Config, es es.AggregateStore) *applyJobHandler {
	return &applyJobHandler{cfg: cfg, es: es}
}

func (c *applyJobHandler) Handle(ctx context.Context, command *ApplyJobCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "applyJobHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	applicationAggregate := aggregate.NewApplicationAggregateWithID(command.AggregateID)
	applicationAggregate.Application = &command.Application
	err := c.es.Exists(ctx, applicationAggregate.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := applicationAggregate.ApplyJob(ctx, command.Application); err != nil {
		return err
	}

	span.LogFields(log.String("application", applicationAggregate.String()))
	return c.es.Save(ctx, applicationAggregate)
}
