package managers

import (
	"cdrateline/auth"
	"cdrateline/database"
	"errors"
	"log"
	"strings"
)

func ChangePassword(email string, oldPassword string, newPassword string, passwordConfirm string) error {
	if newPassword != passwordConfirm {
		log.Println("Passwords do not match")
		return errors.New("passwords do not match")
	}

	if strings.TrimSpace(newPassword) == "" {
		return errors.New("password cannot be blank")
	}

	user, err := database.FetchUserByEmail(email)

	if !auth.CheckPassword(oldPassword, user.Password) {
		return errors.New("old password incorrect")
	}

	if err != nil {
		log.Println(err)
		return err
	}

	newHash, hashErr := auth.HashPassword(newPassword)

	if hashErr != nil {
		log.Println(hashErr)
		return errors.New("could not hash password")
	}

	user.Password = newHash

	updateErr := database.UpdateUser(user)

	if updateErr != nil {
		log.Println(updateErr)
		return errors.New("could not update user")
	}

	return nil
}

func ChangePasswordNoConfirm(email string, newPassword string) error {
	if strings.TrimSpace(newPassword) == "" {
		return errors.New("password cannot be blank")
	}

	user, err := database.FetchUserByEmail(email)
	if err != nil {
		return err
	}

	newHash, hashErr := auth.HashPassword(newPassword)
	if hashErr != nil {
		return hashErr
	}

	user.Password = newHash

	updateErr := database.UpdateUser(user)
	if updateErr != nil {
		return updateErr
	}

	return nil
}
