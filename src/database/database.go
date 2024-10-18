package database

import (
	"time"
	"webdawgengine/config"
	"webdawgengine/crypt"
	"webdawgengine/models"
)

var users []models.User

func Init() {

	defaultPasswd, _ := crypt.HashPassword(config.Config.IdentityDefaultPassword)

	users = []models.User{
		{
			Id:             1,
			Username:       "username",
			Email:          "user@example.com",
			Firstname:      "John",
			Lastname:       "Doe",
			Password:       defaultPasswd,
			FailedAttempts: 0,
			SecurityStamp:  "",
			LastLogin:      time.Now(),
		},
		{
			Id:             2,
			Username:       "example",
			Email:          "example@example.com",
			Firstname:      "Sally",
			Lastname:       "Smith",
			Password:       defaultPasswd,
			FailedAttempts: 0,
			SecurityStamp:  "",
			LastLogin:      time.Now(),
		},
	}
}
