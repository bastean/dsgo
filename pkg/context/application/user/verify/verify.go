package verify

import (
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/domain/types"
)

type Verify struct {
	Repository model.UserRepository
}

func (verify *Verify) Run(id *user.Id) (types.Empty, error) {
	user, err := verify.Repository.Search(&model.UserRepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if user.Verified.Value {
		return nil, nil
	}

	err = verify.Repository.Verify(id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}

func New(repository model.UserRepository) model.UseCase[*user.Id, types.Empty] {
	return &Verify{
		Repository: repository,
	}
}
