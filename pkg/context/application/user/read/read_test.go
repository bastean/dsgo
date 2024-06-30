package read_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/read"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type ReadUseCaseTestSuite struct {
	suite.Suite
	sut        model.UseCase[*user.Id, *user.Primitive]
	repository *persistence.RepositoryMock
}

func (suite *ReadUseCaseTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.sut = read.New(suite.repository)
}

func (suite *ReadUseCaseTestSuite) TestRead() {
	user := user.Random()

	id := user.Id

	criteria := &model.UserRepositorySearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(user)

	expected := user.ToPrimitives()

	actual, err := suite.sut.Run(id)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitReadUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ReadUseCaseTestSuite))
}
