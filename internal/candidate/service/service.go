package service

import (
	"github.com/SEC-Jobstreet/backend-candidate-service/externals"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/commands"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/queries"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/es"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"gorm.io/gorm"
)

type CandidateService struct {
	Commands *commands.CandidateCommands
	Queries  *queries.CandidateQueries
}

func NewCandidateService(config *utils.Config, es es.AggregateStore, db *gorm.DB, jobService *externals.JobServiceGRPC) *CandidateService {
	createProfileHandler := commands.NewCreateProfileHandler(config, es)
	updateProfileHandler := commands.NewUpdateProfileHandler(config, es)
	applyJobHandler := commands.NewApplyJobHandler(config, es)
	saveJobHandler := commands.NewSaveJobHandler(config, es)
	unsaveJobHandler := commands.NewUnsaveJobHandler(config, es)
	candidateCommands := commands.NewCandidateCommands(
		createProfileHandler,
		updateProfileHandler,
		applyJobHandler,
		saveJobHandler,
		unsaveJobHandler,
	)

	getProfileByIDHandler := queries.NewGetProfileByIDHandler(config, es, db)
	getSavedJobListHandler := queries.NewGetSavedJobListHandler(config, es, db, jobService)
	getAppliedCandidateListHandler := queries.NewGetAppliedCandidateListHandler(config, es, db)
	getAppliedCandidateNumberHandler := queries.NewGetAppliedCandidateNumberHandler(config, es, db)
	candidateQueries := queries.NewCandidateQueries(
		getProfileByIDHandler,
		getSavedJobListHandler,
		getAppliedCandidateListHandler,
		getAppliedCandidateNumberHandler,
	)

	return &CandidateService{Commands: candidateCommands, Queries: candidateQueries}
}
