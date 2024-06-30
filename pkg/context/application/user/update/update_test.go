package update_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/update"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/domain/types"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UpdateUseCaseTestSuite struct {
	suite.Suite
	sut        model.UseCase[*user.Primitive, types.Empty]
	repository *persistence.RepositoryMock
}

func (suite *UpdateUseCaseTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.sut = update.New(suite.repository)
}

func (suite *UpdateUseCaseTestSuite) TestUpdate() {
	user := user.Random()

	primitive := user.ToPrimitives()

	criteria := &model.UserRepositorySearchCriteria{
		Id: user.Id,
	}

	suite.repository.On("Search", criteria).Return(user)

	suite.repository.On("Update", user)

	_, err := suite.sut.Run(primitive)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitUpdateUseCaseSuite(t *testing.T) {
	suite.Run(t, new(UpdateUseCaseTestSuite))
}
