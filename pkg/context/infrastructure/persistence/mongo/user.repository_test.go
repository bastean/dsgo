package mongo_test

import (
	"os"
	"testing"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence/mongo"
	"github.com/stretchr/testify/suite"
)

type MongoUserRepositoryTestSuite struct {
	suite.Suite
	sut model.UserRepository
}

func (suite *MongoUserRepositoryTestSuite) SetupTest() {
	uri := os.Getenv("DATABASE_MONGO_URI")

	name := os.Getenv("DATABASE_MONGO_NAME")

	database, _ := mongo.NewMongoDatabase(uri, name)

	collection := "users-test"

	suite.sut, _ = mongo.NewUserCollection(database, collection)
}

func (suite *MongoUserRepositoryTestSuite) TestSave() {
	user := user.Random()
	suite.NoError(suite.sut.Save(user))
}

func (suite *MongoUserRepositoryTestSuite) TestSaveDuplicate() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	err := suite.sut.Save(user)

	var actual *errors.AlreadyExist

	suite.ErrorAs(err, &actual)

	expected := &errors.AlreadyExist{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "HandleMongoDuplicateKeyError",
		What:  "already registered",
		Why: errors.Meta{
			"Field": "Name",
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *MongoUserRepositoryTestSuite) TestUpdate() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Update(user))
}

func (suite *MongoUserRepositoryTestSuite) TestDelete() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Delete(user.Name))
}

func (suite *MongoUserRepositoryTestSuite) TestSearch() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	actual, err := suite.sut.Search(expected.Name)

	suite.NoError(err)

	suite.Equal(expected, actual)
}

func TestIntegrationMongoUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(MongoUserRepositoryTestSuite))
}
