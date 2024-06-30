package persistence

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (repository *RepositoryMock) Save(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *RepositoryMock) Verify(id *user.Id) error {
	repository.Called(id)
	return nil
}

func (repository *RepositoryMock) Update(user *user.User) error {
	repository.Called(user)
	return nil
}

func (repository *RepositoryMock) Delete(id *user.Id) error {
	repository.Called(id)
	return nil
}

func (repository *RepositoryMock) Search(criteria *model.UserRepositorySearchCriteria) (*user.User, error) {
	args := repository.Called(criteria)
	return args.Get(0).(*user.User), nil
}
