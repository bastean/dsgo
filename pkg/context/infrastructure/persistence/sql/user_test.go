package sql_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/repository"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence/sql"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	sut repository.User
}

func (suite *UserTestSuite) SetupTest() {
	var database *sql.Database

	dsn := os.Getenv("DATABASE_MYSQL_DSN")

	name := os.Getenv("DATABASE_MYSQL_NAME")

	database, err := sql.OpenMySQL(dsn, name)

	if err != nil {
		inMemory := "file::memory:?cache=shared"
		database, _ = sql.OpenSQLite(inMemory)
	}

	suite.sut, _ = sql.UserTable(database)
}

func (suite *UserTestSuite) TestSave() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	actual, err := suite.sut.Search(expected.Name)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestSaveErrDuplicatedKey() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	err := suite.sut.Save(user)

	var actual *errors.ErrAlreadyExist

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrAlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Save",
		What:  fmt.Sprintf("%s already registered", user.Name.Value),
		Who:   actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *UserTestSuite) TestUpdate() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	expected.Role = user.RoleWithValidValue()

	suite.NoError(suite.sut.Update(expected))

	actual, err := suite.sut.Search(expected.Name)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestDelete() {
	random := user.Random()

	suite.NoError(suite.sut.Save(random))

	suite.NoError(suite.sut.Delete(random.Name))

	_, err := suite.sut.Search(random.Name)

	suite.Error(err)
}

func (suite *UserTestSuite) TestSearch() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	actual, err := suite.sut.Search(expected.Name)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *UserTestSuite) TestSearchErrRecordNotFound() {
	random := user.Random()

	_, err := suite.sut.Search(random.Name)

	var actual *errors.ErrNotExist

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrNotExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Search",
		What:  fmt.Sprintf("%s not found", random.Name.Value),
		Why: errors.Meta{
			"Index": random.Name.Value,
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}