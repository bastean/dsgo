package create_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/create"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/usecase"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateTestSuite struct {
	suite.Suite
	sut        usecase.Create
	repository *persistence.UserMock
}

func (suite *CreateTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)
	suite.sut = create.New(suite.repository)
}

func (suite *CreateTestSuite) TestCreate() {
	user := user.Random()

	primitive := user.ToPrimitives()

	suite.repository.On("Save", user)

	err := suite.sut.Run(primitive)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
