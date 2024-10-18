package models

import "time"

type User struct {
	Id             int       `db:"id"`
	Username       string    `db:"username"`
	Email          string    `db:"email"`
	Firstname      string    `db:"firstname"`
	Lastname       string    `db:"lastname"`
	Password       string    `db:"password"`
	FailedAttempts int       `db:"failed_attempts"`
	SecurityStamp  string    `db:"security_stamp"`
	LastLogin      time.Time `db:"lastlogin"`
}
