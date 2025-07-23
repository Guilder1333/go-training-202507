package presentation

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler(controller *UserController) http.Handler {
	router := chi.NewRouter()
	router.Get("/users", WithErrorResponse(controller.GetUserByID))
	router.Post("/users", WithErrorResponse(controller.CreateUser))
	router.Delete("/users", WithErrorResponse(controller.DeleteUserById))

	return router
}
