package database

import (
	"log"
	"previous/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

// The database package provides an interface between go code and a relational database.

var db *sqlx.DB

func Init() {
	var err error

	db, err = sqlx.Connect("sqlite3", config.GetConfig().DbConnectionString)

	// allow sqlx scan without all columns present
	db = db.Unsafe()

	if err != nil {
		panic(err.Error())
	}

	// db.SetMaxOpenConns(0)
	// db.SetMaxIdleConns(200)
	// db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Initialized sqlite3 application database connection")
}
