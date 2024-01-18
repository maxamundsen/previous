package data

import (
	"log"
)

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

func AddUser(email string) {
	if email == "" {
		return
	}
	
	sql := "INSERT INTO users (email) VALUES (?)"
	
	stmt, err := db.Prepare(sql)
	
	if err != nil {
		log.Println(err)
	}
	
	_, err = stmt.Exec(email)
	
	if err != nil {
		log.Println(err)
	}
}