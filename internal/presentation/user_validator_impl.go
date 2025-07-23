package presentation

import (
	"encoding/json"
	"errors"
	"fmt"
	"hands_on_go/internal/statuserr"
	"io"
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
		err := errors.New("request query has no id")
		err = statuserr.SetKind(err, statuserr.KindInvalidRequest)
		return UserGetRequest{}, err
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		err = fmt.Errorf("failed to parse user id: %w", err)
		err = statuserr.SetKind(err, statuserr.KindInvalidRequest)
		return UserGetRequest{}, err
	}
	return UserGetRequest{
		ID: id,
	}, nil
}

type createUserBody struct {
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	Age             *int    `json:"age"`
	PhoneNumber     *string `json:"phone"`
	IsPhoneVerified *bool   `json:"phoneVerified"`
}

func (v UserValidatorImpl) ValidateCreateUser(r *http.Request) (*UserCreateRequest, error) {
	var requestBody createUserBody
	err := validateJSON(nil, r.Body, &requestBody)
	err = required(err, requestBody.FirstName, "FirstName")
	err = required(err, requestBody.LastName, "LastName")
	err = required(err, requestBody.Age, "Age")
	err = required(err, requestBody.PhoneNumber, "PhoneNumber")
	err = required(err, requestBody.IsPhoneVerified, "IsPhoneVerified")

	err = stringLength(err, requestBody.FirstName, 0, 100, "FirstName")
	err = stringLength(err, requestBody.LastName, 0, 100, "LastName")
	err = intSize(err, requestBody.Age, 0, 200, "Age")
	err = stringLength(err, requestBody.PhoneNumber, 0, 25, "PhoneNumber")

	if err != nil {
		err = fmt.Errorf("invalid create request body: %w", err)
		err = statuserr.SetKind(err, statuserr.KindInvalidRequest)
		return nil, err
	}

	return &UserCreateRequest{
		FirstName:       *requestBody.FirstName,
		LastName:        *requestBody.LastName,
		Age:             *requestBody.Age,
		PhoneNumber:     *requestBody.PhoneNumber,
		IsPhoneVerified: *requestBody.IsPhoneVerified,
	}, nil
}

func validateJSON(err error, reader io.ReadCloser, result any) error {
	if err != nil {
		return err
	}
	err = json.NewDecoder(reader).Decode(&result)
	if err != nil {
		return fmt.Errorf("json validation failed: %w", err)
	}
	return nil
}

func required[T any](err error, value *T, fieldName string) error {
	if err != nil {
		return err
	}
	if value == nil {
		err = fmt.Errorf("required field '%s' is null", fieldName)
		err = statuserr.SetMessage(err, fmt.Sprintf("required field '%s' is null", fieldName))
		return err
	}
	return nil
}

func stringLength(err error, str *string, min int, max int, fieldName string) error {
	if err != nil || str == nil {
		return err
	}
	v := len(*str)
	if v >= min && v <= max {
		return nil
	}
	err = fmt.Errorf("field '%s' string length does not fit into size constraints [%d, %d]", fieldName, min, max)
	err = statuserr.SetMessage(err, fmt.Sprintf("field '%s' string length does not fit into size constraints [%d, %d]", fieldName, min, max))
	return err
}

func intSize(err error, value *int, min int, max int, fieldName string) error {
	if err != nil || value == nil {
		return err
	}
	v := *value
	if v >= min && v <= max {
		return nil
	}
	err = fmt.Errorf("field '%s' integer value does not fit into size constraints [%d, %d]", fieldName, min, max)
	err = statuserr.SetMessage(err, fmt.Sprintf("field '%s' integer value does not fit into size constraints [%d, %d]", fieldName, min, max))
	return err
}

func (v UserValidatorImpl) ValidateDeleteUserId(r *http.Request) (UserDeleteRequest, error) {
	if !r.URL.Query().Has("id") {
		err := errors.New("delete request query doesn't contain id")
		err = statuserr.SetKind(err, statuserr.KindInvalidRequest)
		return UserDeleteRequest{}, err
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		err = fmt.Errorf("failed to parse user id in delete request: %w", err)
		err = statuserr.SetKind(err, statuserr.KindInvalidRequest)
		return UserDeleteRequest{}, err

	}
	return UserDeleteRequest{
		ID: id,
	}, nil
}
