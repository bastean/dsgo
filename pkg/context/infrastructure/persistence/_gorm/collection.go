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
				"Id": user.Id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Verify(id *user.Id) error {
	filter := bson.D{{Key: "id", Value: id.Value}}

	_, err := db.collection.UpdateOne(context.Background(), filter, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "verified", Value: true},
		}},
	})

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Verify",
			What:  "failure to verify a user",
			Why: errors.Meta{
				"Id": id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Update(user *user.User) error {
	updatedUser := UserDocument(*user.ToPrimitives())

	filter := bson.D{{Key: "id", Value: user.Id.Value}}

	_, err := db.collection.ReplaceOne(context.Background(), filter, &updatedUser)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Update",
			What:  "failure to update a user",
			Why: errors.Meta{
				"Id": user.Id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Delete(id *user.Id) error {
	filter := bson.D{{Key: "id", Value: id.Value}}

	_, err := db.collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Delete",
			What:  "failure to delete a user",
			Why: errors.Meta{
				"Id": id.Value,
			},
			Who: err,
		})
	}

	return nil
}

func (db *UserCollection) Search(criteria *model.UserRepositorySearchCriteria) (*user.User, error) {
	var filter bson.D
	var index string

	switch {
	case criteria.Id != nil:
		filter = bson.D{{Key: "id", Value: criteria.Id.Value}}
		index = criteria.Id.Value
	case criteria.Email != nil:
		filter = bson.D{{Key: "email", Value: criteria.Email.Value}}
		index = criteria.Email.Value
	}

	result := db.collection.FindOne(context.Background(), filter)

	if err := result.Err(); err != nil {
		return nil, HandleMongoDocumentNotFound(index, err)
	}

	primitive := new(user.Primitive)

	err := result.Decode(primitive)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "Search",
			What:  "failure to decode a result",
			Why: errors.Meta{
				"Index": index,
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
				"Index":     index,
			},
			Who: err,
		})
	}

	return user, nil
}

func NewMongoCollection(mdb *MongoDB, collectionName string) (model.UserRepository, error) {
	collection := mdb.Database.Collection(collectionName)

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
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
