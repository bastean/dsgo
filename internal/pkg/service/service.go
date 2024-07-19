package service

import (
	"github.com/bastean/dsgo/internal/pkg/service/env"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/bastean/dsgo/internal/pkg/service/user"
	"github.com/bastean/dsgo/pkg/context/domain/repository"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence/mysql"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence/sqlite"
)

var (
	err    error
	MySQL  *mysql.MySQL
	SQLite *sqlite.SQLite
)

var (
	Service = &struct {
		MySQL, SQLite string
	}{
		MySQL:  log.Service("mysql"),
		SQLite: log.Service("sqlite"),
	}
	Module = &struct {
		User string
	}{
		User: log.Module("user"),
	}
)

func OpenMySQL() error {
	MySQL, err = mysql.Open(
		env.Database.MySQL.DSN,
		env.Database.MySQL.Name,
	)

	if err != nil {
		return errors.BubbleUp(err, "OpenMySQL")
	}

	return nil
}

func OpenSQLite() error {
	SQLite, err = sqlite.Open(
		env.Database.SQLite.Name,
	)

	if err != nil {
		return errors.BubbleUp(err, "OpenSQLite")
	}

	return nil
}

func StartModuleUser() error {
	var table repository.User

	switch {
	case MySQL != nil:
		table, err = mysql.UserTable(MySQL)
	case SQLite != nil:
		table, err = sqlite.UserTable(SQLite)
	}

	if err != nil {
		return errors.BubbleUp(err, "StartModuleUser")
	}

	user.Start(
		table,
	)

	return nil
}

func Run() error {
	log.EstablishingConnectionWith(Service.MySQL)

	if err = OpenMySQL(); err != nil {
		log.ConnectionFailedWith(Service.MySQL)

		log.Error(err.Error())

		log.EstablishingConnectionWith(Service.SQLite)

		if err = OpenSQLite(); err != nil {
			return errors.BubbleUp(err, "Run")
		}

		log.ConnectionEstablishedWith(Service.SQLite)
	} else {
		log.ConnectionEstablishedWith(Service.MySQL)
	}

	log.Starting(Module.User)

	err = StartModuleUser()

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	log.Started(Module.User)

	return nil
}

func CloseMySQL() error {
	err = mysql.Close(MySQL)

	if err != nil {
		return errors.BubbleUp(err, "CloseMySQL")
	}

	return nil
}

func CloseSQLite() error {
	err = sqlite.Close(SQLite)

	if err != nil {
		return errors.BubbleUp(err, "CloseSQLite")
	}

	return nil
}

func Stop() error {
	switch {
	case MySQL != nil:
		log.ClosingConnectionWith(Service.MySQL)
		err = mysql.Close(MySQL)
	case SQLite != nil:
		log.ClosingConnectionWith(Service.SQLite)
		err = sqlite.Close(SQLite)
	}

	if err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	return nil
}
