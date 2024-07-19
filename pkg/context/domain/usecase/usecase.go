package usecase

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
)

type (
	Create interface {
		Run(*user.Primitive) error
	}
	Read interface {
		Run(name string) (*user.Primitive, error)
	}
	Update interface {
		Run(*user.Primitive) error
	}
	Delete interface {
		Run(name string) error
	}
)
