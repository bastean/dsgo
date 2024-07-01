package update_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/update"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UpdateUseCaseTestSuite struct {
	suite.Suite
	sut        *update.Update
	repository *persistence.RepositoryMock
}

func (suite *UpdateUseCaseTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.sut = update.New(suite.repository)
}

func (suite *UpdateUseCaseTestSuite) TestUpdate() {
	user := user.Random()

	primitive := user.ToPrimitives()

	suite.repository.On("Update", user)

	err := suite.sut.Run(primitive)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitUpdateUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UpdateUseCaseTestSuite))
}
