package user

import (
	"github.com/bastean/dsgo/pkg/context/shared/domain/errors"
	"github.com/bastean/dsgo/pkg/context/shared/infrastructure/persistences"
	"github.com/bastean/dsgo/pkg/context/user/domain/model"
	"github.com/bastean/dsgo/pkg/context/user/infrastructure/persistence"
)

func MongoCollection(database *persistences.MongoDB, name string, hashing model.Hashing) (model.Repository, error) {
	collection, err := persistence.NewMongoCollection(database, name, hashing)

	if err != nil {
		return nil, errors.BubbleUp(err, "MongoCollection")
	}

	return collection, nil
}
