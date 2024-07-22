package service

import (
	"github.com/bastean/dsgo/internal/pkg/service/env"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/bastean/dsgo/internal/pkg/service/user"
	"github.com/bastean/dsgo/pkg/context/infrastructure/persistence/sql"
)

var (
	err      error
	Database *sql.Database
)

var (
	Service = &struct {
		Database, MySQL, SQLite string
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
	Database, err = sql.OpenMySQL(
		env.Database.MySQL.DSN,
		env.Database.MySQL.Name,
	)

	if err != nil {
		return errors.BubbleUp(err, "OpenMySQL")
	}

	return nil
}

func OpenSQLite() error {
	Database, err = sql.OpenSQLite(
		env.Database.SQLite.Name,
	)

	if err != nil {
		return errors.BubbleUp(err, "OpenSQLite")
	}

	return nil
}

func StartModuleUser() error {
	table, err := sql.UserTable(Database)

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

		Service.Database = Service.SQLite
	} else {
		log.ConnectionEstablishedWith(Service.MySQL)
		Service.Database = Service.MySQL
	}

	log.Starting(Module.User)

	err = StartModuleUser()

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	log.Started(Module.User)

	return nil
}

func CloseDatabase() error {
	log.ClosingConnectionWith(Service.Database)

	err = sql.Close(Database)

	if err != nil {
		return errors.BubbleUp(err, "CloseDatabase")
	}

	log.ConnectionClosedWith(Service.Database)

	return nil
}

func Stop() error {
	if err := CloseDatabase(); err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	return nil
}
