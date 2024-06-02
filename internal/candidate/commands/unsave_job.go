package commands

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/aggregate"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

type UnsaveJobCommandHandler interface {
	Handle(ctx context.Context, command *UnSaveJobCommand) error
}

type unsaveJobHandler struct {
	cfg *utils.Config
	es  es.AggregateStore
}

func NewUnsaveJobHandler(cfg *utils.Config, es es.AggregateStore) *unsaveJobHandler {
	return &unsaveJobHandler{cfg: cfg, es: es}
}

func (c *unsaveJobHandler) Handle(ctx context.Context, command *UnSaveJobCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "unsaveJobHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	savedJobAggregate, err := aggregate.LoadSavedJobAggregate(ctx, c.es, command.SavedJob.CandidateID+"-"+command.SavedJob.JobID.String())
	if err != nil {
		return err
	}

	if err := savedJobAggregate.UnsaveJob(ctx, command.SavedJob); err != nil {
		return err
	}

	span.LogFields(log.String("unsavedJob", savedJobAggregate.String()))
	return c.es.Save(ctx, savedJobAggregate)
}
