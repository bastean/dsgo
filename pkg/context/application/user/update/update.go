package update

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
)

type Update struct {
	Repository model.UserRepository
}

func (update *Update) Run(primitive *user.Primitive) error {
	user, err := user.New(primitive)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = update.Repository.Update(user)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func New(repository model.UserRepository) *Update {
	return &Update{
		Repository: repository,
	}
}
