package delete

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
)

type Delete struct {
	Repository model.UserRepository
}

func (delete *Delete) Run(name string) error {
	nameVO, err := user.NewName(name)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	user, err := delete.Repository.Search(nameVO)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = delete.Repository.Delete(user.Name)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func New(repository model.UserRepository) *Delete {
	return &Delete{
		Repository: repository,
	}
}
