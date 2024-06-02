package queries

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"gorm.io/gorm"
)

type GetSavedJobListQueryHandler interface {
	Handle(ctx context.Context, command *GetSavedJobListQuery) ([]models.SavedJob, error)
}

type getSavedJobListHandler struct {
	cfg *utils.Config
	es  es.AggregateStore
	db  *gorm.DB
}

func NewGetSavedJobListHandler(cfg *utils.Config, es es.AggregateStore, db *gorm.DB) *getSavedJobListHandler {
	return &getSavedJobListHandler{cfg: cfg, es: es, db: db}
}

func (q *getSavedJobListHandler) Handle(ctx context.Context, query *GetSavedJobListQuery) ([]models.SavedJob, error) {
	savedJobList := []models.SavedJob{}
	err := q.db.Where("candidate_id = ?", query.Username).Find(&savedJobList).Error
	if err != nil {
		return nil, err
	}

	return savedJobList, nil
}
