package presentation

import (
	"errors"
	"net/http"
)

var ErrInvalidGetRequest = errors.New("get user request is invalid")

type UserGetRequest struct {
	ID int
}

type UserValidator interface {
	ValidateGetUserId(r *http.Request) (UserGetRequest, error)
}
