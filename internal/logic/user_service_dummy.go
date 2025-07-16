package logic

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
