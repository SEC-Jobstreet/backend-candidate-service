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

type SaveJobCommandHandler interface {
	Handle(ctx context.Context, command *SaveJobCommand) error
}

type saveJobHandler struct {
	cfg *utils.Config
	es  es.AggregateStore
}

func NewSaveJobHandler(cfg *utils.Config, es es.AggregateStore) *saveJobHandler {
	return &saveJobHandler{cfg: cfg, es: es}
}

func (c *saveJobHandler) Handle(ctx context.Context, command *SaveJobCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "saveJobHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	savedJobAggregate := aggregate.NewSavedJobAggregateWithID(command.AggregateID)
	savedJobAggregate.SavedJob = &command.SavedJob
	err := c.es.Exists(ctx, savedJobAggregate.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := savedJobAggregate.SaveJob(ctx, command.SavedJob); err != nil {
		return err
	}

	span.LogFields(log.String("savedJob", savedJobAggregate.String()))
	return c.es.Save(ctx, savedJobAggregate)
}
