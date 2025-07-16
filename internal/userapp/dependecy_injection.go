package userapp

import (
	"hands_on_go/internal/logic"
	"hands_on_go/internal/presentation"
	"net/http"

	"github.com/rs/zerolog/log"
)

type application struct {
	rootHandler http.Handler
}

func newApplication() (*application, error) {
	controller, err := presentation.NewUserController(
		presentation.NewUserValidatorImpl(),
		logic.NewUserServiceDummy(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create user controller")
	}

	return &application{
		rootHandler: presentation.NewHandler(controller),
	}, nil
}

func closeApplication(app *application) {
	// TODO
}
