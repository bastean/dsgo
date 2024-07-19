package env

import (
	"os"
)

type MySQL struct {
	DSN, Name string
}

type SQLite struct {
	Name string
}

var Database = &struct {
	*MySQL
	*SQLite
}{
	MySQL: &MySQL{
		DSN:  os.Getenv("DATABASE_MYSQL_DSN"),
		Name: os.Getenv("DATABASE_MYSQL_NAME"),
	},
	SQLite: &SQLite{
		Name: os.Getenv("DATABASE_SQLITE_NAME"),
	},
}
