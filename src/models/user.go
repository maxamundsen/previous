package models

import "time"

type User struct {
	Id             int
	Username       string
	Email          string
	Firstname      string
	Lastname       string
	Password       string
	FailedAttempts int
	SecurityStamp  string
	LastLogin      time.Time

	PermissionAdmin bool
}
