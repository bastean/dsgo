package mysql_test

import (
	"os"
	"testing"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence/mysql"
	"github.com/stretchr/testify/suite"
)

type MySQLUserRepositoryTestSuite struct {
	suite.Suite
	sut model.UserRepository
}

func (suite *MySQLUserRepositoryTestSuite) SetupTest() {
	dsn := os.Getenv("DATABASE_MYSQL_DSN")

	name := os.Getenv("DATABASE_MYSQL_NAME")

	database, _ := mysql.NewMySQLDatabase(dsn, name)

	suite.sut, _ = mysql.NewUserTable(database)
}

func (suite *MySQLUserRepositoryTestSuite) TestSave() {
	user := user.Random()
	suite.NoError(suite.sut.Save(user))
}

func (suite *MySQLUserRepositoryTestSuite) TestSaveDuplicate() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	err := suite.sut.Save(user)

	var actual *errors.AlreadyExist

	suite.ErrorAs(err, &actual)

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Save",
		What:  "already registered",
		Who:   actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *MySQLUserRepositoryTestSuite) TestUpdate() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Update(user))
}

func (suite *MySQLUserRepositoryTestSuite) TestDelete() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Delete(user.Name))
}

func (suite *MySQLUserRepositoryTestSuite) TestSearch() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	actual, err := suite.sut.Search(expected.Name)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func TestIntegrationMySQLUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(MySQLUserRepositoryTestSuite))
}
