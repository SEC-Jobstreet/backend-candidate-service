package main

import (
	"context"
	"github.com/SEC-Jobstreet/backend-application-service/api/handlers"
	"github.com/SEC-Jobstreet/backend-application-service/api/services"
	"os"
	"os/signal"
	"syscall"

	server "github.com/SEC-Jobstreet/backend-application-service/api"
	"github.com/SEC-Jobstreet/backend-application-service/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var interruptSignals = []os.Signal{os.Interrupt, syscall.SIGINT, syscall.SIGTERM}

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)

	runGinServer(ctx, waitGroup, config)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runGinServer(ctx context.Context, waitGroup *errgroup.Group, config utils.Config) {
	// repos

	// services
	profileService := services.NewProfileService()

	// handlers
	profileHandler := handlers.NewProfileHandler(profileService)

	// inject
	ginServer, err := server.NewServer(
		config,
		profileHandler,
	)

	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	err = ginServer.Start(ctx, waitGroup, config.ListenIP+":"+config.ListenPort)
	if err != nil {
		log.Fatal().Msg("cannot start server")
	}
}
