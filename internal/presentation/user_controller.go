package presentation

import (
	"encoding/json"
	"errors"
	"hands_on_go/internal/logic"
	"net/http"

	"github.com/rs/zerolog/log"
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

func (u *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// get ID from request validate the id
	requestInfo, err := u.getRequestValidator.ValidateGetUserId(r)
	if errors.Is(err, ErrInvalidGetRequest) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		log.Warn().Err(err).Msg("Invalid user get request")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		log.Error().Err(err).Msg("Unexpected error while handling user get request")
		return
	}

	// pass id to business logic (get user)
	user, err := u.userService.GetUserByID(requestInfo.ID)
	if errors.Is(err, logic.ErrUserNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		log.Warn().Err(err).Msg("User not found for get request")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		log.Error().Err(err).Msg("Unexpected error while handling user get request")
		return
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		log.Error().Err(err).Msg("failed to serialize response for user get request")
		return
	}

	w.Write(responseStr)
	w.WriteHeader(http.StatusOK)
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse and validate request body.
	requestInfo, err := u.getRequestValidator.ValidateCreateUser(r)
	if errors.Is(err, ErrInvalidCreateRequest) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		log.Warn().Err(err).Msg("Invalid user create request")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		log.Error().Err(err).Msg("Unexpected error while handling user create request")
		return
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
	if errors.Is(err, logic.ErrUserCreationFailed) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user creation failed"))
		log.Warn().Err(err).Msg("Failed to create user")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		log.Error().Err(err).Msg("Unexpected error while handling user create request")
		return
	}
	// Serialize response and write it.

	response := createUserResponse{
		ID: userId,
	}
	responseStr, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("something went wrong"))
		log.Error().Err(err).Msg("failed to serialize response for user create request")
		return
	}

	w.Write(responseStr)
	w.WriteHeader(http.StatusCreated)
}
