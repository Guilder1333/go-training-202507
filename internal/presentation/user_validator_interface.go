package presentation

import (
	"net/http"
)

// var ErrInvalidGetRequest = errors.New("get user request is invalid")
// var ErrInvalidCreateRequest = errors.New("create user request is invalid")
// var ErrInvalidDeleteRequest = errors.New("detele user request is invalid")

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

type UserDeleteRequest struct {
	ID int
}

type UserValidator interface {
	ValidateGetUserId(r *http.Request) (UserGetRequest, error)
	ValidateCreateUser(r *http.Request) (*UserCreateRequest, error)
	ValidateDeleteUserId(r *http.Request) (UserDeleteRequest, error)
}
