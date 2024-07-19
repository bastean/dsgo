package update

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/repository"
	"github.com/bastean/dsgo/pkg/context/domain/usecase"
)

type Update struct {
	repository.User
}

func (update *Update) Run(primitive *user.Primitive) error {
	new, err := user.New(primitive)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	_, err = update.User.Search(new.Name)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = update.User.Update(new)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func New(repository repository.User) usecase.Update {
	return &Update{
		User: repository,
	}
}
