package service

import (
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

func NewCandidateService(
	config *utils.Config,
	es es.AggregateStore,
	db *gorm.DB,
) *CandidateService {

	createProfileHandler := commands.NewCreateProfileHandler(config, es)
	updateProfileHandler := commands.NewUpdateProfileHandler(config, es)

	getProfileByIDHandler := queries.NewGetProfileByIDHandler(config, es, db)

	candidateCommands := commands.NewCandidateCommands(
		createProfileHandler,
		updateProfileHandler,
	)
	candidateQueries := queries.NewCandidateQueries(getProfileByIDHandler)

	return &CandidateService{Commands: candidateCommands, Queries: candidateQueries}
}
