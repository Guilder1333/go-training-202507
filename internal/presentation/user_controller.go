package presentation

import (
	"encoding/json"
	"errors"
	"fmt"
	"hands_on_go/internal/logic"
	"net/http"
)

type UserController struct {
	getRequestValidator UserValidator
	userService         logic.UserService
}

type getUserByIdResponse struct {
	ID              int    `json:"-"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Age             int    `json:"age"`
	PhoneNumber     string `json:"phone"`
	IsPhoneVerified bool   `json:"phoneVerified"`
}

type createUserResponse struct {
	ID int `json:"id"`
}

func NewUserController(getRequestValidator UserValidator, userService logic.UserService) (*UserController, error) {
	if getRequestValidator == nil {
		return nil, errors.New("getRequestValidator can't be null")
	}
	if userService == nil {
		return nil, errors.New("userService can't be null")
	}

	return &UserController{
		getRequestValidator: getRequestValidator,
		userService:         userService,
	}, nil
}

func (u *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) error {
	// get ID from request validate the id
	requestInfo, err := u.getRequestValidator.ValidateGetUserId(r)
	if err != nil {
		return fmt.Errorf("user validation for get request failed: %w", err)
	}

	// pass id to business logic (get user)
	user, err := u.userService.GetUserByID(requestInfo.ID)
	if err != nil {
		return fmt.Errorf("get user by id failed for get request: %w", err)
	}

	// map user to json struct
	userResponse := getUserByIdResponse{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Age:             user.Age,
		PhoneNumber:     user.PhoneNumber,
		IsPhoneVerified: user.IsPhoneVerified,
	}

	// serialize json struct to string and send it as a response.
	responseStr, err := json.Marshal(userResponse)
	if err != nil {
		return fmt.Errorf("serializing response failed for get request: %w", err)
	}

	w.Write(responseStr)
	w.WriteHeader(http.StatusOK)
	return nil
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) error {
	// Parse and validate request body.
	requestInfo, err := u.getRequestValidator.ValidateCreateUser(r)
	if err != nil {
		return fmt.Errorf("failed to validate create user request: %w", err)
	}
	// Pass data to business logic
	user := logic.User{
		ID:              0,
		FirstName:       requestInfo.FirstName,
		LastName:        requestInfo.LastName,
		Age:             requestInfo.Age,
		PhoneNumber:     requestInfo.PhoneNumber,
		IsPhoneVerified: requestInfo.IsPhoneVerified,
	}

	userId, err := u.userService.Create(&user)
	if err != nil {
		return fmt.Errorf("failed to create new user: %w", err)
	}

	// Serialize response and write it.

	response := createUserResponse{
		ID: userId,
	}
	responseStr, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to serialize response for create user: %w", err)
	}

	w.Write(responseStr)
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (u *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) error {
	// get ID from request validate the id
	requestInfo, err := u.getRequestValidator.ValidateDeleteUserId(r)
	if err != nil {
		return fmt.Errorf("failed to validate delete user request: %w", err)
	}

	// pass id to business logic (delete user)
	err = u.userService.Delete(requestInfo.ID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
