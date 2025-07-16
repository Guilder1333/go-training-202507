package presentation

import (
	"errors"
	"net/http"
)

var ErrInvalidGetRequest = errors.New("get user request is invalid")
var ErrInvalidCreateRequest = errors.New("create user request is invalid")

type UserGetRequest struct {
	ID int
}

type UserCreateRequest struct {
	FirstName       string
	LastName        string
	Age             int
	PhoneNumber     string
	IsPhoneVerified bool
}

type UserValidator interface {
	ValidateGetUserId(r *http.Request) (UserGetRequest, error)
	ValidateCreateUser(r *http.Request) (*UserCreateRequest, error)
}
