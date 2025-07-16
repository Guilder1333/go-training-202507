package logic

import "math/rand"

type UserServiceDummy struct {
}

func NewUserServiceDummy() UserServiceDummy {
	return UserServiceDummy{}
}

func (s UserServiceDummy) GetUserByID(userId int) (*User, error) {
	if userId == 404 {
		return nil, ErrUserNotFound
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
		return 0, ErrUserCreationFailed
	}

	return rand.Intn(1000) + 1, nil
}

func (s UserServiceDummy) Delete(userId int) error {
	if userId == 404 {
		return ErrUserNotFound
	}

	return nil
}
