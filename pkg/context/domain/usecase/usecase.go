package usecase

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
)

type (
	Create interface {
		Run(*user.Primitive) error
	}
	Read interface {
		Run(*user.Name) (*user.Primitive, error)
	}
	Update interface {
		Run(*user.Primitive) error
	}
	Delete interface {
		Run(*user.Name) error
	}
)
