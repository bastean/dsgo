package read

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/repository"
	"github.com/bastean/dsgo/pkg/context/domain/usecase"
)

type Read struct {
	repository.User
}

func (read *Read) Run(name string) (*user.Primitive, error) {
	nameVO, err := user.NewName(name)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	user, err := read.User.Search(nameVO)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return user.ToPrimitives(), nil
}

func New(repository repository.User) usecase.Read {
	return &Read{
		User: repository,
	}
}
