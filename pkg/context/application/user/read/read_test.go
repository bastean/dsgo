package read_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/read"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/usecase"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type ReadTestSuite struct {
	suite.Suite
	sut        usecase.Read
	repository *persistence.UserMock
}

func (suite *ReadTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)
	suite.sut = read.New(suite.repository)
}

func (suite *ReadTestSuite) TestRead() {
	user := user.Random()

	suite.repository.On("Search", user.Name).Return(user)

	expected := user.ToPrimitives()

	actual, err := suite.sut.Run(user.Name.Value)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitReadSuite(t *testing.T) {
	suite.Run(t, new(ReadTestSuite))
}
