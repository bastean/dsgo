package persistence

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (repository *UserRepositoryMock) Save(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *UserRepositoryMock) Update(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *UserRepositoryMock) Delete(name *user.Name) error {
	repository.Called(name)
	return nil
}

func (repository *UserRepositoryMock) Search(name *user.Name) (*user.User, error) {
	args := repository.Called(name)
	return args.Get(0).(*user.User), nil
}
