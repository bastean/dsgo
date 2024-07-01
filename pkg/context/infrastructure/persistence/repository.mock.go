package persistence

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (repository *RepositoryMock) Save(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *RepositoryMock) Update(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *RepositoryMock) Delete(name *user.Name) error {
	repository.Called(name)
	return nil
}

func (repository *RepositoryMock) Search(name *user.Name) (*user.User, error) {
	args := repository.Called(name)
	return args.Get(0).(*user.User), nil
}
