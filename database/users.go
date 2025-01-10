package database

import (
	"previous/.jet/model"

	. "previous/.jet/table"
	. "github.com/go-jet/jet/v2/sqlite"
)

func FetchUserById(userid int32) (model.User, error) {
	dest := model.User{}

	stmt := SELECT(
		User.AllColumns,
	).FROM(User).WHERE(
		User.ID.EQ(Int32(int32(userid))),
	)

	err := stmt.Query(db, &dest)

	return dest, err
}

func FetchUserByUsername(username string) (model.User, error) {
	dest := model.User{}

	stmt := SELECT(
		User.AllColumns,
	).FROM(
		User,
	).WHERE(
		User.Username.EQ(String(username)),
	)

	err := stmt.Query(db, &dest)
	return dest, err
}

func FetchUserSecurityStamp(userid int) (string, error) {
	stamp := ""

	stmt := SELECT(
		User.SecurityStamp,
	).FROM(User).WHERE(
		User.ID.EQ(Int32(int32(userid))),
	)

	err := stmt.Query(db, &stamp)
	return stamp, err
}

func UpdateUser(user model.User) error {
	stmt := User.UPDATE(User.MutableColumns).
	MODEL(user).
	WHERE(User.ID.EQ(Int32(user.ID)))

	_, err := stmt.Exec(db)
	return err
}
