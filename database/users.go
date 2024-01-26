package database

import (
	"log"
)

type User struct {
	Id    int
	Email string
}

func FetchUsers() []User {
	users := make([]User, 0)

	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)

	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
	}

	count := 0

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Email)

		users = append(users, user)
		count += 1
	}

	log.Printf("Fetched %d users\n", count)

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

	log.Println("User added. Email: " + email)
}

func DeleteAllUsers() {
	sql := "TRUNCATE TABLE users"

	stmt, err := db.Prepare(sql)

	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec()

	if err != nil {
		log.Println(err)
	}

	log.Println("Deleted all users")
}
