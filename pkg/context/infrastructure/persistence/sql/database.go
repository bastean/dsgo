package sql

import (
	"fmt"

	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	Session *gorm.DB
}

var Config = &gorm.Config{
	TranslateError: true,
}

func OpenMySQL(dsn, name string) (*Database, error) {
	queryParams := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf("%s/%s?%s", dsn, name, queryParams)

	session, err := gorm.Open(mysql.Open(dsn), Config)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "OpenMySQL",
			What:  "Failure connecting to mysql",
			Who:   err,
		})
	}

	return &Database{
		Session: session,
	}, nil
}

func OpenSQLite(filename string) (*Database, error) {
	session, err := gorm.Open(sqlite.Open(filename), Config)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "OpenSQLite",
			What:  "Failure connecting to sqlite",
			Who:   err,
		})
	}

	return &Database{
		Session: session,
	}, nil
}

func Close(database *Database) error {
	session, err := database.Session.DB()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "Failure to obtain database",
			Who:   err,
		})
	}

	err = session.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "Failure to close connection",
			Who:   err,
		})
	}

	return nil
}
