package queries

import (
	"context"

	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"gorm.io/gorm"
)

type GetAppliedCandidateListQueryHandler interface {
	Handle(ctx context.Context, command *GetAppliedCandidateListQuery) ([]models.Profile, int64, error)
}

type getAppliedCandidateListHandler struct {
	cfg *utils.Config
	es  es.AggregateStore
	db  *gorm.DB
}

func NewGetAppliedCandidateListHandler(cfg *utils.Config, es es.AggregateStore, db *gorm.DB) *getAppliedCandidateListHandler {
	return &getAppliedCandidateListHandler{cfg: cfg, es: es, db: db}
}

func (q *getAppliedCandidateListHandler) Handle(ctx context.Context, query *GetAppliedCandidateListQuery) ([]models.Profile, int64, error) {
	profileList := []models.Profile{}
	tx := q.db.Model(&models.Application{}).Where("job_id = ?", query.JobID)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	var total int64
	tx.Count(&total)

	list := []models.Application{}
	err := tx.Limit(query.PageSize).Offset((query.PageID - 1) * query.PageSize).Find(&list).Error
	if err != nil {
		return nil, 0, tx.Error
	}

	for _, application := range list {
		candidate := models.Profile{
			Username: application.CandidateID,
		}
		if er := q.db.First(&candidate).Error; er != nil {
			return nil, 0, tx.Error
		}
		profileList = append(profileList, candidate)
	}

	return profileList, total, nil
}
