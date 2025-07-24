package logic

type UserRepository interface {
	CheckExists(id int) (bool, error)
	DeleteByID(id int) error

	GetUser(id int) (*User, error)
	CreateUser(user *User) (int, error)
}
