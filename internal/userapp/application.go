package userapp

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

func Run() {
	// Setup Application
	application, err := newApplication()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create new application")
	}

	// Run HTTP Server
	runServer(application)

	// Shutdown HTTP server
	// Teardown application
	log.Info().Msg("Performing tear down of application")
	closeApplication(application)
}

func runServer(application *application) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGTERM, syscall.SIGINT)
	errchan := make(chan error, 1)

	server := http.Server{
		Addr:    ":8080",
		Handler: application.rootHandler,
	}

	go serverServe(errchan, &server)
	log.Info().Msg("Listening for port 8080")

	select {
	case sig := <-sigchan:
		log.Info().Msg("Got termination signal from os " + sig.String())

		err := server.Shutdown(context.TODO())
		if err != nil {
			log.Error().Err(err).Msg("Server shutdown failed.")
		}
	case err := <-errchan:
		log.Error().Err(err).Msg("Failed to start listening for port")
	}
}

func serverServe(errchan chan error, server *http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		errchan <- err
	}
}
