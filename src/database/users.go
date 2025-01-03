package database

import (
	"errors"
	"previous/models"
)

func FetchUserById(userid int) (models.User, error) {
	for _, v := range users {
		if v.Id == userid {
			return v, nil
		}
	}

	return models.User{}, errors.New("no user found")
}

func FetchUserByUsername(username string) (models.User, error) {
	for _, v := range users {
		if v.Username == username {
			return v, nil
		}
	}

	return models.User{}, errors.New("no user found")
}

func FetchUserSecurityStamp(userid int) (string, error) {
	for _, v := range users {
		if v.Id == userid {
			return v.SecurityStamp, nil
		}
	}

	return "", errors.New("no user found")
}

func UpdateUser(user models.User) error {
	for _, v := range users {
		if v.Id == user.Id {
			v = user
			return nil
		}
	}

	return errors.New("could not update user. no user found")
}
