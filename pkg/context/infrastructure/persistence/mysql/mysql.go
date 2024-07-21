package mysql

import (
	"fmt"

	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Session *gorm.DB
}

func Open(dsn, name string) (*MySQL, error) {
	queryParams := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf("%s/%s?%s", dsn, name, queryParams)

	session, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "NewMySQLDatabase",
			What:  "Failure connecting to mysql",
			Who:   err,
		})
	}

	return &MySQL{
		Session: session,
	}, nil
}

func Close(mySQL *MySQL) error {
	session, err := mySQL.Session.DB()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "CloseMySQLDatabase",
			What:  "Failure to obtain mysql database",
			Who:   err,
		})
	}

	err = session.Close()

	if err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "CloseMySQLDatabase",
			What:  "Failure to close connection with mysql",
			Who:   err,
		})
	}

	return nil
}
