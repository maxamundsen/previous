package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitializeDb(connectionString string) {
	var err error

	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}

	log.Println("Initialized MySQL application database connection")
}
