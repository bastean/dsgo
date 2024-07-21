package sqlite

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	Session *gorm.DB
}

func Open(filename string) (*SQLite, error) {
	session, err := gorm.Open(sqlite.Open(filename), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "New",
			What:  "Failure connecting to sqlite",
			Who:   err,
		})
	}

	return &SQLite{
		Session: session,
	}, nil
}

func Close(sqLite *SQLite) error {
	session, err := sqLite.Session.DB()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "Failure to obtain sqlite database",
			Who:   err,
		})
	}

	err = session.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "Failure to close connection with sqlite",
			Who:   err,
		})
	}

	return nil
}
