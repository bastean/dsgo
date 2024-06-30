package verify_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/application/user/verify"
	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/domain/types"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type VerifyHandlerTestSuite struct {
	suite.Suite
	sut        model.UseCase[*user.Id, types.Empty]
	repository *persistence.RepositoryMock
}

func (suite *VerifyHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.sut = verify.New(suite.repository)
}

func (suite *VerifyHandlerTestSuite) TestVerify() {
	user := user.Random()

	id := user.Id

	criteria := &model.UserRepositorySearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(user)

	suite.repository.On("Verify", id)

	_, err := suite.sut.Run(id)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitVerifyHandlerSuite(t *testing.T) {
	suite.Run(t, new(VerifyHandlerTestSuite))
}
