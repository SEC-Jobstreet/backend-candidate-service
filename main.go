package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"syscall"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	_ "github.com/SEC-Jobstreet/backend-candidate-service/docs"
	"github.com/SEC-Jobstreet/backend-candidate-service/externals"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/models"
	postgres_projection "github.com/SEC-Jobstreet/backend-candidate-service/internal/candidate/projection"
	"github.com/SEC-Jobstreet/backend-candidate-service/internal/server"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/eventstroredb"
	"github.com/SEC-Jobstreet/backend-candidate-service/pkg/logger"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

//	@title			candidatee Service API
//	@version		1.0
//	@description	This is a candidate Service Server.

// @host		localhost:4000
// @BasePath	/api/v1
func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	sqlDB, err := sql.Open("pgx", config.DBSource)
	if err != nil {
		log.Fatal().Msg("cannot connect to db")
	}

	store, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("cannot connect to db")
	}

	err = models.MigrateDB(store)
	if err != nil {
		log.Fatal().Msg("could not migrate db")
	}

	awsHandler := externals.NewAWSHandler()

	appLogger := logger.NewAppLogger(&logger.Config{
		LogLevel: "debug",
		DevMode:  false,
		Encoder:  "console",
	})
	appLogger.InitLogger()
	appLogger.WithName("candidateService")

	db, err := eventstroredb.NewEventStoreDB(eventstroredb.EventStoreConfig{
		ConnectionString: config.EventStoreConnectionString,
	})
	if err != nil {
		log.Fatal().Msg("could not connect to eventstore db")
	}
	defer db.Close()

	transportOptions := grpc.WithTransportCredentials(insecure.NewCredentials())
	gRPCconn, err := grpc.Dial(config.JobServiceGRPCAddress, transportOptions)
	if err != nil {
		log.Fatal().Msg("cannot dial grpc job server")
	}

	postgresProjection := postgres_projection.NewCandidateProjection(appLogger, db, store, &config)

	go func() {
		err := postgresProjection.Subscribe(ctx, []string{""}, 60, postgresProjection.ProcessEvents)
		if err != nil {
			log.Fatal().Msg("(candidateProjection.Subscribe)")
			stop()
		}
	}()

	waitGroup, ctx := errgroup.WithContext(ctx)

	runGinServer(ctx, waitGroup, appLogger, config, db, store, awsHandler, gRPCconn)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runGinServer(ctx context.Context, waitGroup *errgroup.Group, appLogger logger.Logger, config utils.Config, db *esdb.Client, store *gorm.DB, awsHandler *externals.AWSHandler, gRPCconn *grpc.ClientConn) {
	ginServer, err := server.NewServer(config, appLogger, db, store, awsHandler, gRPCconn)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	err = ginServer.Start(ctx, waitGroup, config.RESTfulServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start server")
	}
}
