package read

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
)

type Read struct {
	Repository model.UserRepository
}

func (read *Read) Run(id *user.Id) (*user.Primitive, error) {
	user, err := read.Repository.Search(&model.UserRepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return user.ToPrimitives(), nil
}

func New(repository model.UserRepository) model.UseCase[*user.Id, *user.Primitive] {
	return &Read{
		Repository: repository,
	}
}
