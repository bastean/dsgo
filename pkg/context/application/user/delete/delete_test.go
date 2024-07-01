package delete_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/delete"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type DeleteUseCaseTestSuite struct {
	suite.Suite
	sut        *delete.Delete
	repository *persistence.RepositoryMock
}

func (suite *DeleteUseCaseTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.sut = delete.New(suite.repository)
}

func (suite *DeleteUseCaseTestSuite) TestDelete() {
	user := user.Random()

	name := user.Name

	suite.repository.On("Search", name).Return(user)

	suite.repository.On("Delete", name)

	err := suite.sut.Run(name.Value)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitDeleteUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DeleteUseCaseTestSuite))
}
