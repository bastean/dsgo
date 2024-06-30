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

type MongoRepositoryTestSuite struct {
	suite.Suite
	sut model.UserRepository
}

func (suite *MongoRepositoryTestSuite) SetupTest() {
	uri := os.Getenv("DATABASE_URI")

	databaseName := "dsgo-test"

	database, _ := mongo.NewMongoDatabase(uri, databaseName)

	collectionName := "users-test"

	suite.sut, _ = mongo.NewMongoCollection(database, collectionName)
}

func (suite *MongoRepositoryTestSuite) TestSave() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))
}

func (suite *MongoRepositoryTestSuite) TestSaveDuplicate() {
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
			"Field": "Id",
		},
		Who: actual.Who,
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *MongoRepositoryTestSuite) TestVerify() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Verify(user.Id))
}

func (suite *MongoRepositoryTestSuite) TestUpdate() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Update(user))
}

func (suite *MongoRepositoryTestSuite) TestDelete() {
	user := user.Random()

	suite.NoError(suite.sut.Save(user))

	suite.NoError(suite.sut.Delete(user.Id))
}

func (suite *MongoRepositoryTestSuite) TestSearch() {
	expected := user.Random()

	suite.NoError(suite.sut.Save(expected))

	criteria := &model.UserRepositorySearchCriteria{
		Id: expected.Id,
	}

	user, err := suite.sut.Search(criteria)

	suite.NoError(err)

	actual := user

	suite.Equal(expected, actual)
}

func TestIntegrationMongoRepositorySuite(t *testing.T) {
	suite.Run(t, new(MongoRepositoryTestSuite))
}
