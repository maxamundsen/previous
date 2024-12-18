package database

import (
	"saral/crypt"
)

type Migration struct {
	PreviousHashBase64 string
	SQL string
}

var Migrations []Migration

func AppendMigration(m Migration) {
	version := len(Migrations)

	if version > 0 {
		previousMigration := Migrations[version - 1]

		m.PreviousHashBase64, _ = crypt.HighwayHash(previousMigration.SQL)
	}

	Migrations = append(Migrations, m)
}

func InitializeMigrations() {
	AppendMigration(Migration{
		SQL: "CREATE TABLE TEST",
	})

	AppendMigration(Migration{
		SQL: "CREATE TABLE TEST2",
	})
}

func GetLatestMigration() Migration {
	return Migrations[len(Migrations) - 1]
}

func ConcatAllMigrations() string {
	out := ""

	for _, v := range Migrations {
		out += v.SQL + " "
	}

	return out
}

func MigrateUp() error {
	return nil
}

func GetLatestMigrationFromConnectedDB() Migration {
	return Migration{}
}