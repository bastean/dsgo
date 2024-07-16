package persistence

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (repository *UserMock) Save(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *UserMock) Update(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *UserMock) Delete(name *user.Name) error {
	repository.Called(name)
	return nil
}

func (repository *UserMock) Search(name *user.Name) (*user.User, error) {
	args := repository.Called(name)
	return args.Get(0).(*user.User), nil
}
