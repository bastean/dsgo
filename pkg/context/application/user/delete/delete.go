package delete

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/repository"
	"github.com/bastean/dsgo/pkg/context/domain/usecase"
)

type Delete struct {
	repository.User
}

func (delete *Delete) Run(name string) error {
	nameVO, err := user.NewName(name)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	user, err := delete.User.Search(nameVO)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = delete.User.Delete(user.Name)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func New(repository repository.User) usecase.Delete {
	return &Delete{
		User: repository,
	}
}
