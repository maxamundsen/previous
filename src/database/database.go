package database

import (
	"time"
	"saral/config"
	"saral/crypt"
	"saral/models"
)

var users []models.User

func Init() {
	adminPass, _ := crypt.HashPassword("admin")
	pass, _ := crypt.HashPassword(config.GetConfig().IdentityDefaultPassword)

	// default users, no database so hardcode this into the database "init"
	users = []models.User{
		{
			Id:              1,
			Username:        "admin",
			Email:           "user@example.com",
			Firstname:       "John",
			Lastname:        "Doe",
			Password:        adminPass,
			FailedAttempts:  0,
			SecurityStamp:   "",
			LastLogin:       time.Now(),
			PermissionAdmin: true,
		},
		{
			Id:              2,
			Username:        "example",
			Email:           "example@example.com",
			Firstname:       "Sally",
			Lastname:        "Smith",
			Password:        pass,
			FailedAttempts:  0,
			SecurityStamp:   "",
			LastLogin:       time.Now(),
			PermissionAdmin: false,
		},
	}
}
