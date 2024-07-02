package read_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/read"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type ReadUseCaseTestSuite struct {
	suite.Suite
	sut        *read.Read
	repository *persistence.UserRepositoryMock
}

func (suite *ReadUseCaseTestSuite) SetupTest() {
	suite.repository = new(persistence.UserRepositoryMock)
	suite.sut = read.New(suite.repository)
}

func (suite *ReadUseCaseTestSuite) TestRead() {
	user := user.Random()

	name := user.Name

	suite.repository.On("Search", name).Return(user)

	expected := user.ToPrimitives()

	actual, err := suite.sut.Run(name.Value)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitReadUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ReadUseCaseTestSuite))
}
