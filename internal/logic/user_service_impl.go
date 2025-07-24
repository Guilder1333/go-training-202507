package logic

import (
	"errors"
	"fmt"
	"hands_on_go/internal/statuserr"
)

type UserServiceImpl struct {
	repository UserRepository
}

func NewUserServiceImpl(repository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (s *UserServiceImpl) GetUserByID(userId int) (*User, error) {
	ok, err := s.repository.CheckExists(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists: %w", err)
	}
	if !ok {
		err = errors.New("user doesn't exist")
		err = statuserr.SetKind(err, statuserr.KindUserNotFound)
		return nil, err
	}

	user, err := s.repository.GetUser(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserServiceImpl) Create(user *User) (int, error) {
	id, err := s.repository.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}

func (s *UserServiceImpl) Delete(userId int) error {
	ok, err := s.repository.CheckExists(userId)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	if !ok {
		err = errors.New("user for deletion doesn't exist")
		err = statuserr.SetKind(err, statuserr.KindUserNotFound)
		return err
	}
	err = s.repository.DeleteByID(userId)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
