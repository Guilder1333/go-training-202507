package logic

// var ErrUserNotFound = errors.New("failed to find user")
// var ErrUserCreationFailed = errors.New("failed to create new user")

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
	Create(*User) (int, error)
	Delete(userId int) error
}
