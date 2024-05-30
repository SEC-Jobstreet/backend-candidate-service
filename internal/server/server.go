package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/SEC-Jobstreet/backend-candidate-service/externals"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/api"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/service"
	es "github.com/SEC-Jobstreet/backend-candidate-service/pkg/es/store"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/logger"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	log    logger.Logger
	config utils.Config
	store  *gorm.DB
	cs     *service.CandidateService
	media  *externals.AWSHandler
	db     *esdb.Client
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config utils.Config, appLogger logger.Logger, db *esdb.Client, store *gorm.DB, awsHandler *externals.AWSHandler) (*Server, error) {
	awsHandler.Init(config)

	aggregateStore := es.NewAggregateStore(appLogger, db)

	server := &Server{
		config: config,
		store:  store,
		cs:     service.NewCandidateService(&config, aggregateStore, store),
		media:  awsHandler,
		db:     db,
	}

	return server, nil
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(ctx context.Context, waitGroup *errgroup.Group, address string) error {

	ginRouter := gin.Default()

	candidateHandlers := api.NewCandidateHandlers(server.config, server.cs, ginRouter, server.media)
	router := candidateHandlers.SetupRouter()

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	var err error
	waitGroup.Go(func() error {
		log.Info().Msgf("RESTFUL API server serve at %s", address)
		err = srv.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Fatal().Msg("RESTFUL API server failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown RESTFUL API server")

		err = srv.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to shutdown RESTFUL API server")
			return err
		}
		log.Info().Msg("RESTFUL API server is stopped")
		return nil
	})

	return err
}
