package model

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
)

type UserRepositorySearchCriteria struct {
	*user.Id
	*user.Email
}

type UserRepository interface {
	Save(user *user.User) error
	Verify(id *user.Id) error
	Update(user *user.User) error
	Delete(id *user.Id) error
	Search(criteria *UserRepositorySearchCriteria) (*user.User, error)
}
