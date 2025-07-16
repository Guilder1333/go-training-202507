package presentation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler(controller *UserController) http.Handler {
	router := chi.NewRouter()
	router.Get("/users", controller.GetUserByID)
	router.Post("/users", controller.CreateUser)

	return router
}
