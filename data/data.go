package data

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitializeDb(connectionString string) {
	_, err := sql.Open("mysql", connectionString)
	
	if err != nil {
		log.Println(err.Error())
	}
}



type User struct {
	Id int
	Email string
}

func FetchUsers() []User {
	users := make([]User, 0)
	
	rows, err := db.Query("SELECT * FROM users")
	
	defer rows.Close()

    if err != nil {
        log.Println(err.Error())
    }
    
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email)

		users = append(users, user)
	}	
	
	return users
}