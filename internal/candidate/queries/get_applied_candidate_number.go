package queries

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"gorm.io/gorm"
)

type GetAppliedCandidateNumberQueryHandler interface {
	Handle(ctx context.Context, command *GetAppliedCandidateNumberQuery) (int64, error)
}

type getAppliedCandidateNumberHandler struct {
	cfg *utils.Config
	es  es.AggregateStore
	db  *gorm.DB
}

func NewGetAppliedCandidateNumberHandler(cfg *utils.Config, es es.AggregateStore, db *gorm.DB) *getAppliedCandidateNumberHandler {
	return &getAppliedCandidateNumberHandler{cfg: cfg, es: es, db: db}
}

func (q *getAppliedCandidateNumberHandler) Handle(ctx context.Context, query *GetAppliedCandidateNumberQuery) (int64, error) {

	var count int64
	err := q.db.Model(&models.Application{}).Where("job_id = ?", query.JobID).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}
