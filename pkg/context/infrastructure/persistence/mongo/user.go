package mongo

import (
	"context"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDocument struct {
	Name string `bson:"name,omitempty"`
	Role string `bson:"role,omitempty"`
}

type UserCollection struct {
	collection *mongo.Collection
}

func (db *UserCollection) Save(user *user.User) error {
	newUser := UserDocument(*user.ToPrimitives())

	_, err := db.collection.InsertOne(context.Background(), &newUser)

	if mongo.IsDuplicateKeyError(err) {
		return errors.BubbleUp(HandleMongoDuplicateKeyError(err), "Save")
	}

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Save",
			What:  "failure to save a user",
			Why: errors.Meta{
				"Name": user.Name.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Update(user *user.User) error {
	updatedUser := UserDocument(*user.ToPrimitives())

	filter := bson.D{{Key: "name", Value: user.Name.Value}}

	_, err := db.collection.ReplaceOne(context.Background(), filter, &updatedUser)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "failure to update a user",
			Why: errors.Meta{
				"Name": user.Name.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Delete(name *user.Name) error {
	filter := bson.D{{Key: "name", Value: name.Value}}

	_, err := db.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "failure to delete a user",
			Why: errors.Meta{
				"Name": name.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Search(name *user.Name) (*user.User, error) {
	filter := bson.D{{Key: "name", Value: name.Value}}

	result := db.collection.FindOne(context.Background(), filter)

	if err := result.Err(); err != nil {
		return nil, HandleMongoDocumentNotFound(name.Value, err)
	}

	primitive := new(user.Primitive)

	err := result.Decode(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "failure to decode a result",
			Why: errors.Meta{
				"Index": name.Value,
			},
			Who: err,
		})
	}

	user, err := user.FromPrimitives(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "failure to create an aggregate from a primitive",
			Why: errors.Meta{
				"Primitive": primitive,
				"Index":     name.Value,
			},
			Who: err,
		})
	}

	return user, nil
}

func NewUserCollection(mdb *MongoDB, collectionName string) (model.UserRepository, error) {
	collection := mdb.Database.Collection(collectionName)

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewMongoCollection",
			What:  "failure to create indexes for user collection",
			Why: errors.Meta{
				"Collection": collectionName,
			},
			Who: err,
		})
	}

	return &UserCollection{
		collection: collection,
	}, nil
}
