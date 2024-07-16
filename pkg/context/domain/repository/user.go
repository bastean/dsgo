package repository

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
)

type User interface {
	Save(*user.User) error
	Update(*user.User) error
	Delete(*user.Name) error
	Search(*user.Name) (*user.User, error)
}
