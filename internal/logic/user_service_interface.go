package logic

import "errors"

var ErrUserNotFound = errors.New("failed to find user")

type User struct {
	ID              int
	FirstName       string
	LastName        string
	Age             int
	PhoneNumber     string
	IsPhoneVerified bool
}

type UserService interface {
	GetUserByID(userId int) (*User, error)
}
