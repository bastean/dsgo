package model

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
)

type UserRepository interface {
	Save(user *user.User) error
	Update(user *user.User) error
	Delete(name *user.Name) error
	Search(name *user.Name) (*user.User, error)
}
