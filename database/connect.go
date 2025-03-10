package database

import (
	"previous/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// The database package provides an interface between go code and a relational database.

var DB *sqlx.DB

func Init() {
	var err error

	DB, err = sqlx.Connect("sqlite3", config.GetConfig().DbConnectionString)

	if err != nil {
		panic(err.Error())
	}

	DB = DB.Unsafe()

	// db.SetMaxOpenConns(0)
	// db.SetMaxIdleConns(200)
	// db.SetConnMaxLifetime(5 * time.Minute)
}
