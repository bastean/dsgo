package sqlite

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	Client *gorm.DB
}

func New(filename string) (*SQLite, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "New",
			What:  "failure connecting to sqlite",
			Who:   err,
		})
	}

	return &SQLite{
		Client: db,
	}, nil
}

func Close(sdb *SQLite) error {
	db, err := sdb.Client.DB()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "failure to obtain sqlite database",
			Who:   err,
		})
	}

	err = db.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "failure to close connection with sqlite",
			Who:   err,
		})
	}

	return nil
}
