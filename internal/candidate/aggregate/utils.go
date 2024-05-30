package aggregate

import (
	"context"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

// GetProfileAggregateID get profile aggregate id for eventstoredb
func GetProfileAggregateID(eventAggregateID string) string {
	return strings.ReplaceAll(eventAggregateID, "profile-", "")
}

func IsAggregateNotFound(aggregate es.Aggregate) bool {
	return aggregate.GetVersion() == 0
}

func LoadProfileAggregate(ctx context.Context, eventStore es.AggregateStore, aggregateID string) (*ProfileAggregate, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoadProfileAggregate")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", aggregateID))

	profile := NewProfileAggregateWithID(aggregateID)

	err := eventStore.Exists(ctx, profile.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return nil, err
	}

	if err := eventStore.Load(ctx, profile); err != nil {
		return nil, err
	}

	return profile, nil
}
