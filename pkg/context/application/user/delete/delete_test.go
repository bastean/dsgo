package delete_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/delete"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/usecase"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type DeleteTestSuite struct {
	suite.Suite
	sut        usecase.Delete
	repository *persistence.UserMock
}

func (suite *DeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)
	suite.sut = delete.New(suite.repository)
}

func (suite *DeleteTestSuite) TestDelete() {
	user := user.Random()

	suite.repository.On("Delete", user.Name)

	err := suite.sut.Run(user.Name)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitDeleteUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
