package queries

import (
	"context"
	"fmt"

	"github.com/SEC-Jobstreet/backend-candidate-service/externals"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"gorm.io/gorm"
)

type GetSavedJobListQueryHandler interface {
	Handle(ctx context.Context, command *GetSavedJobListQuery) ([]models.Jobs, error)
}

type getSavedJobListHandler struct {
	cfg            *utils.Config
	es             es.AggregateStore
	db             *gorm.DB
	jobServiceGRPC *externals.JobServiceGRPC
}

func NewGetSavedJobListHandler(cfg *utils.Config, es es.AggregateStore, db *gorm.DB, jobService *externals.JobServiceGRPC) *getSavedJobListHandler {
	return &getSavedJobListHandler{cfg: cfg, es: es, db: db, jobServiceGRPC: jobService}
}

func (q *getSavedJobListHandler) Handle(ctx context.Context, query *GetSavedJobListQuery) ([]models.Jobs, error) {
	savedJobList := []models.SavedJob{}
	err := q.db.Where("candidate_id = ?", query.Username).Find(&savedJobList).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(savedJobList)

	if len(savedJobList) > 0 {
		var jobs []models.Jobs
		for _, savedJob := range savedJobList {
			job, err := q.jobServiceGRPC.GetJob(savedJob.JobID.String())
			if err == nil {
				fmt.Println(job)
				jobs = append(jobs, *job)
			}
		}

		return jobs, nil
	}
	return []models.Jobs{}, nil
}
