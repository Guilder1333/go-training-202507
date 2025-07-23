package logic

import (
	"errors"
	"hands_on_go/internal/statuserr"
	"math/rand"
)

type UserServiceDummy struct {
}

func NewUserServiceDummy() UserServiceDummy {
	return UserServiceDummy{}
}

func (s UserServiceDummy) GetUserByID(userId int) (*User, error) {
	if userId == 404 {
		err := errors.New("user not found")
		err = statuserr.SetKind(err, statuserr.KindUserNotFound)
		return nil, err
	}

	return &User{
		ID:              userId,
		FirstName:       "Jane",
		LastName:        "Doe",
		Age:             45,
		PhoneNumber:     "12356436456",
		IsPhoneVerified: true,
	}, nil
}

func (s UserServiceDummy) Create(user *User) (int, error) {
	if user.FirstName == "invalid" {
		err := errors.New("create user failed")
		err = statuserr.SetKind(err, statuserr.KindCreateUserFailed)
		return 0, err
	}

	return rand.Intn(1000) + 1, nil
}

func (s UserServiceDummy) Delete(userId int) error {
	if userId == 404 {
		err := errors.New("user not found")
		err = statuserr.SetKind(err, statuserr.KindUserNotFound)
		return err
	}

	return nil
}
