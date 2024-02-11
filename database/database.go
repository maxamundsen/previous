package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"webdawgengine/config"
)

var db *sql.DB

func InitializeDb() {
	var err error

	db, err = sql.Open("mysql", config.GetConnectionString())

	if err != nil {
		panic(err.Error())
	}

	log.Println("Initialized MySQL application database connection")
}
