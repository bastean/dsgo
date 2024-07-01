package create

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
)

type Create struct {
	Repository model.UserRepository
}

func (create *Create) Run(primitive *user.Primitive) error {
	user, err := user.New(primitive)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = create.Repository.Save(user)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func New(repository model.UserRepository) *Create {
	return &Create{
		Repository: repository,
	}
}
