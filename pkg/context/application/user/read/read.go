package read

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
)

type Read struct {
	Repository model.UserRepository
}

func (read *Read) Run(name string) (*user.Primitive, error) {
	nameVO, err := user.NewName(name)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	user, err := read.Repository.Search(nameVO)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return user.ToPrimitives(), nil
}

func New(repository model.UserRepository) *Read {
	return &Read{
		Repository: repository,
	}
}
