package logic

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceImpl_Delete_CheckExistsFailed_ReturnError(t *testing.T) {
	checkExistsError := errors.New("Check exists failed mock")

	repostory := userRepositoryMock{
		checkExists: func(id int) (bool, error) {
			return false, checkExistsError
		},
	}
	userService := NewUserServiceImpl(&repostory)

	err := userService.Delete(124345)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, checkExistsError)
}

func TestUserServiceImpl_Delete_Success(t *testing.T) {
	repostory := userRepositoryMock{
		checkExists: func(id int) (bool, error) { return true, nil },
		deleteById: func(id int) error {
			return nil
		},
	}
	userService := NewUserServiceImpl(&repostory)

	err := userService.Delete(123)

	assert.Nil(t, err)
}

// And two more successful tests for Create and Get
