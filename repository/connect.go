package repository

import (
	"log"
	"previous/config"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// The database package provides an interface between go code and a relational database.

var db *sql.DB

func Init() {
	var err error

	db, err = sql.Open("sqlite3", config.GetConfig().DbConnectionString)

	if err != nil {
		panic(err.Error())
	}

	// db.SetMaxOpenConns(0)
	// db.SetMaxIdleConns(200)
	// db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Initialized sqlite3 application database connection")
}
