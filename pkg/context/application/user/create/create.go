package create

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/domain/types"
)

type Create struct {
	Repository model.UserRepository
}

func (create *Create) Run(primitive *user.Primitive) (types.Empty, error) {
	user, err := user.New(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = create.Repository.Save(user)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

func New(repository model.UserRepository) model.UseCase[*user.Primitive, types.Empty] {
	return &Create{
		Repository: repository,
	}
}
