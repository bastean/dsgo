package delete_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/shared/domain/models"
	"github.com/bastean/dsgo/pkg/context/shared/domain/types"
	"github.com/bastean/dsgo/pkg/context/user/application/delete"
	"github.com/bastean/dsgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/dsgo/pkg/context/user/domain/model"
	"github.com/bastean/dsgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/dsgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type DeleteHandlerTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*delete.Command]
	usecase    models.UseCase[*delete.Input, types.Empty]
	hashing    *cryptographic.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *DeleteHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.usecase = &delete.Delete{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &delete.Handler{
		UseCase: suite.usecase,
	}
}

func (suite *DeleteHandlerTestSuite) TestDelete() {
	user := aggregate.RandomUser()

	command := &delete.Command{
		Id:       user.Id.Value(),
		Password: user.Password.Value(),
	}

	criteria := &model.RepositorySearchCriteria{
		Id: user.Id,
	}

	suite.repository.On("Search", criteria).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value(), user.Password.Value()).Return(false)

	suite.repository.On("Delete", user.Id)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitDeleteHandlerSuite(t *testing.T) {
	suite.Run(t, new(DeleteHandlerTestSuite))
}
