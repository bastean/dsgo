package create_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/create"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateUseCaseTestSuite struct {
	suite.Suite
	sut        *create.Create
	repository *persistence.UserRepositoryMock
}

func (suite *CreateUseCaseTestSuite) SetupTest() {
	suite.repository = new(persistence.UserRepositoryMock)
	suite.sut = create.New(suite.repository)
}

func (suite *CreateUseCaseTestSuite) TestCreate() {
	user := user.Random()

	primitive := user.ToPrimitives()

	suite.repository.On("Save", user)

	err := suite.sut.Run(primitive)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitCreateUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CreateUseCaseTestSuite))
}
