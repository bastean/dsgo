package delete_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/delete"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/domain/types"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type DeleteUseCaseTestSuite struct {
	suite.Suite
	sut        model.UseCase[*user.Id, types.Empty]
	repository *persistence.RepositoryMock
}

func (suite *DeleteUseCaseTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.sut = delete.New(suite.repository)
}

func (suite *DeleteUseCaseTestSuite) TestDelete() {
	user := user.Random()

	id := user.Id

	criteria := &model.UserRepositorySearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(user)

	suite.repository.On("Delete", id)

	_, err := suite.sut.Run(id)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitDeleteUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DeleteUseCaseTestSuite))
}
