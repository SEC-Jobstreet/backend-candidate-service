package queries

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"gorm.io/gorm"
)

type GetProfileByIDQueryHandler interface {
	Handle(ctx context.Context, command *GetProfileByIDQuery) (*models.Profile, error)
}

type getProfileByIDHandler struct {
	cfg *utils.Config
	es  es.AggregateStore
	db  *gorm.DB
}

func NewGetProfileByIDHandler(cfg *utils.Config, es es.AggregateStore, db *gorm.DB) *getProfileByIDHandler {
	return &getProfileByIDHandler{cfg: cfg, es: es, db: db}
}

func (q *getProfileByIDHandler) Handle(ctx context.Context, query *GetProfileByIDQuery) (*models.Profile, error) {
	candidate := &models.Profile{
		Username: query.ID,
	}
	err := q.db.First(candidate).Error
	if err != nil {
		return nil, err
	}

	return candidate, nil
}
