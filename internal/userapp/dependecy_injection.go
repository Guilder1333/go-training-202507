package userapp

import (
	"hands_on_go/internal/presentation"
	"net/http"
)

type application struct {
	rootHandler http.Handler
}

func newApplication() (*application, error) {
	return &application{
		rootHandler: presentation.NewHandler(),
	}, nil
}

func closeApplication(app *application) {
	// TODO
}
