package create_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/create"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/domain/types"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateUseCaseTestSuite struct {
	suite.Suite
	sut        model.UseCase[*user.Primitive, types.Empty]
	repository *persistence.RepositoryMock
}

func (suite *CreateUseCaseTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.sut = create.New(suite.repository)
}

func (suite *CreateUseCaseTestSuite) TestCreate() {
	user := user.Random()

	primitive := user.ToPrimitives()

	suite.repository.On("Save", user)

	_, err := suite.sut.Run(primitive)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitCreateUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CreateUseCaseTestSuite))
}
