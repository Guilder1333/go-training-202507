package logic


type userRepositoryMock struct {
	checkExists func(id int) (bool, error)
	deleteById  func(id int) error
	getUser     func(id int) (*User, error)
	createUser  func(user *User) (int, error)
}

func (r *userRepositoryMock) CheckExists(id int) (bool, error) {
	return r.checkExists(id)
}
func (r *userRepositoryMock) DeleteByID(id int) error {
	return r.deleteById(id)
}
func (r *userRepositoryMock) GetUser(id int) (*User, error) {
	return r.getUser(id)
}
func (r *userRepositoryMock) CreateUser(user *User) (int, error) {
	return r.createUser(user)
}

