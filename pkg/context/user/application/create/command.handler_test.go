package create_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/shared/domain/models"
	"github.com/bastean/dsgo/pkg/context/shared/domain/types"
	"github.com/bastean/dsgo/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/dsgo/pkg/context/user/application/create"
	"github.com/bastean/dsgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/dsgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateHandlerTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*create.Command]
	usecase    models.UseCase[*aggregate.User, types.Empty]
	repository *persistence.RepositoryMock
	broker     *communications.BrokerMock
}

func (suite *CreateHandlerTestSuite) SetupTest() {
	suite.broker = new(communications.BrokerMock)

	suite.repository = new(persistence.RepositoryMock)

	suite.usecase = &create.Create{
		Repository: suite.repository,
	}

	suite.sut = &create.Handler{
		UseCase: suite.usecase,
		Broker:  suite.broker,
	}
}

func (suite *CreateHandlerTestSuite) TestCreate() {
	command := create.RandomCommand()

	user, _ := aggregate.NewUser(&aggregate.UserPrimitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
	})

	messages := user.Messages

	suite.repository.On("Save", user)

	suite.broker.On("PublishMessages", messages)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitCreateHandlerSuite(t *testing.T) {
	suite.Run(t, new(CreateHandlerTestSuite))
}
