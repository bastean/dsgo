package mysql_test

import (
	"os"
	"testing"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/repository"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence/mysql"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	sut repository.User
}

func (suite *UserTestSuite) SetupTest() {
	dsn := os.Getenv("DATABASE_MYSQL_DSN")

	name := os.Getenv("DATABASE_MYSQL_NAME")

	database, _ := mysql.Open(dsn, name)

	suite.sut, _ = mysql.UserTable(database)
}

func (suite *UserTestSuite) TestSave() {
	user := user.Random()
	suite.NoError(suite.sut.Save(user))
}

func (suite *UserTestSuite) TestSaveDuplicate() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	err := suite.sut.Save(user)

	var actual *errors.ErrAlreadyExist

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrAlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Save",
		What:  "already registered",
		Who:   actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *UserTestSuite) TestUpdate() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Update(user))
}

func (suite *UserTestSuite) TestDelete() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Delete(user.Name))
}

func (suite *UserTestSuite) TestSearch() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	actual, err := suite.sut.Search(expected.Name)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func TestIntegrationUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
