package delete

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/domain/types"
)

type Delete struct {
	Repository model.UserRepository
}

func (delete *Delete) Run(id *user.Id) (types.Empty, error) {
	user, err := delete.Repository.Search(&model.UserRepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = delete.Repository.Delete(user.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

func New(repository model.UserRepository) model.UseCase[*user.Id, types.Empty] {
	return &Delete{
		Repository: repository,
	}
}
