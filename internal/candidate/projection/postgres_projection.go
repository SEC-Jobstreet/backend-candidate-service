package postgres_projection

import (
	"context"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/events"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/constants"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/logger"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/tracing"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

const (
	ProjectionGroupName = "candidate"
)

type postgresProjection struct {
	log          logger.Logger
	db           *esdb.Client
	cfg          *utils.Config
	postgresRepo *gorm.DB
}

func NewCandidateProjection(log logger.Logger, db *esdb.Client, postgresRepo *gorm.DB, cfg *utils.Config) *postgresProjection {
	return &postgresProjection{log: log, db: db, postgresRepo: postgresRepo, cfg: cfg}
}

type Worker func(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error

func (o *postgresProjection) Subscribe(ctx context.Context, prefixes []string, poolSize int, worker Worker) error {
	o.log.Infof("(starting candidate subscription) prefixes: {%+v}", prefixes)

	err := o.db.CreatePersistentSubscriptionAll(ctx, ProjectionGroupName, esdb.PersistentAllSubscriptionOptions{
		Filter: &esdb.SubscriptionFilter{Type: esdb.StreamFilterType, Prefixes: prefixes},
	})
	if err != nil {
		if subscriptionError, ok := err.(*esdb.PersistentSubscriptionError); !ok || ok && (subscriptionError.Code != 6) {
			o.log.Errorf("(CreatePersistentSubscriptionAll) err: {%v}", subscriptionError.Error())
		}
	}

	stream, err := o.db.ConnectToPersistentSubscription(
		ctx,
		constants.EsAll,
		ProjectionGroupName,
		esdb.ConnectToPersistentSubscriptionOptions{},
	)
	if err != nil {
		return err
	}
	defer stream.Close()

	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i <= poolSize; i++ {
		g.Go(o.runWorker(ctx, worker, stream, i))
	}
	return g.Wait()
}

func (o *postgresProjection) runWorker(ctx context.Context, worker Worker, stream *esdb.PersistentSubscription, i int) func() error {
	return func() error {
		return worker(ctx, stream, i)
	}
}

func (o *postgresProjection) ProcessEvents(ctx context.Context, stream *esdb.PersistentSubscription, workerID int) error {

	for {
		event := stream.Recv()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if event.SubscriptionDropped != nil {
			o.log.Errorf("(SubscriptionDropped) err: {%v}", event.SubscriptionDropped.Error)
			return errors.Wrap(event.SubscriptionDropped.Error, "Subscription Dropped")
		}

		if event.EventAppeared != nil {
			o.log.ProjectionEvent(constants.PostgresProjection, ProjectionGroupName, event.EventAppeared, workerID)

			err := o.When(ctx, es.NewEventFromRecorded(event.EventAppeared.Event))
			if err != nil {
				o.log.Errorf("(postgresProjection.when) err: {%v}", err)

				if err := stream.Nack(err.Error(), esdb.Nack_Retry, event.EventAppeared); err != nil {
					o.log.Errorf("(stream.Nack) err: {%v}", err)
					return errors.Wrap(err, "stream.Nack")
				}
			}

			err = stream.Ack(event.EventAppeared)
			if err != nil {
				o.log.Errorf("(stream.Ack) err: {%v}", err)
				return errors.Wrap(err, "stream.Ack")
			}
			o.log.Infof("(ACK) event commit: {%v}", *event.EventAppeared.Commit)
		}
	}
}

func (o *postgresProjection) When(ctx context.Context, evt es.Event) error {
	ctx, span := tracing.StartProjectionTracerSpan(ctx, "postgresProjection.When", evt)
	defer span.Finish()
	span.LogFields(log.String("AggregateID", evt.GetAggregateID()), log.String("EventType", evt.GetEventType()))

	switch evt.GetEventType() {

	case events.ProfileCreated:
		return o.onProfileCreate(ctx, evt)
	case events.ProfileUpdated:
		return o.onProfileUpdate(ctx, evt)
	default:
		o.log.Warnf("(postgresProjection) [When unknown EventType] eventType: {%s}", evt.EventType)
		return es.ErrInvalidEventType
	}
}
