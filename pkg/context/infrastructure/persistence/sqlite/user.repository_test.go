package sqlite_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence/sqlite"
	"github.com/stretchr/testify/suite"
)

type SQLiteUserRepositoryTestSuite struct {
	suite.Suite
	sut     model.UserRepository
	dirTemp string
}

func (suite *SQLiteUserRepositoryTestSuite) SetupTest() {
	suite.dirTemp = "temp"

	os.Mkdir(suite.dirTemp, os.ModePerm)

	filename := os.Getenv("SQLITE_DATABASE")

	path := filepath.Join(suite.dirTemp, filename)

	database, _ := sqlite.New(path)

	suite.sut, _ = sqlite.NewUserTable(database)
}

func (suite *SQLiteUserRepositoryTestSuite) TestSave() {
	user := user.Random()
	suite.NoError(suite.sut.Save(user))
}

func (suite *SQLiteUserRepositoryTestSuite) TestSaveDuplicate() {
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

func (suite *SQLiteUserRepositoryTestSuite) TestUpdate() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Update(user))
}

func (suite *SQLiteUserRepositoryTestSuite) TestDelete() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Delete(user.Name))
}

func (suite *SQLiteUserRepositoryTestSuite) TestSearch() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	actual, err := suite.sut.Search(expected.Name)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func (suite *SQLiteUserRepositoryTestSuite) TearDownTest() {
	os.RemoveAll(suite.dirTemp)
}

func TestIntegrationSQLiteUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(SQLiteUserRepositoryTestSuite))
}
