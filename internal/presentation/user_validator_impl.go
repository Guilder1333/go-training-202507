package presentation

import (
	"fmt"
	"net/http"
	"strconv"
)

type UserValidatorImpl struct {
}

func NewUserValidatorImpl() UserValidatorImpl {
	return UserValidatorImpl{}
}

func (v UserValidatorImpl) ValidateGetUserId(r *http.Request) (UserGetRequest, error) {
	if !r.URL.Query().Has("id") {
		return UserGetRequest{}, ErrInvalidGetRequest
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return UserGetRequest{},
			fmt.Errorf("%w: failed to parse user id: %w", ErrInvalidGetRequest, err)
	}
	return UserGetRequest{
		ID: id,
	}, nil
}
