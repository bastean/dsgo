package env

import (
	"os"
)

var (
	DatabaseMySQLDSN  = os.Getenv("DATABASE_MYSQL_DSN")
	DatabaseMySQLName = os.Getenv("DATABASE_MYSQL_NAME")
)

var (
	DatabaseSQLiteName = os.Getenv("DATABASE_SQLITE_NAME")
)
