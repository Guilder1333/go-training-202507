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

// TODO:
// 1. stop using the commented errors in those two files
// 2. Use SetKind to make new errors
// 3. Use GetKind to get information about error
// 4. User WithErrorResponse for all request handlers
