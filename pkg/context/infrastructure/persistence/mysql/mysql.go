package mysql

import (
	"fmt"

	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Client *gorm.DB
}

func NewMySQLDatabase(dsn, name string) (*MySQL, error) {
	queryParams := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf("%s/%s?%s", dsn, name, queryParams)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewMySQLDatabase",
			What:  "failure connecting to mysql",
			Who:   err,
		})
	}

	return &MySQL{
		Client: db,
	}, nil
}

func CloseMySQLDatabase(mdb *MySQL) error {
	db, err := mdb.Client.DB()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "CloseMySQLDatabase",
			What:  "failure to obtain mysql database",
			Who:   err,
		})
	}

	err = db.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "CloseMySQLDatabase",
			What:  "failure to close connection with mysql",
			Who:   err,
		})
	}

	return nil
}
