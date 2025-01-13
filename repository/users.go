package repository

import (
	"database/sql"
	"previous/.jet/model"

	. "previous/.jet/table"

	. "github.com/go-jet/jet/v2/sqlite"
)

type UserRepository struct{}

func (u UserRepository) FetchById(id int32) (model.User, error) {
	user := model.User{}

	stmt := SELECT(
		User.AllColumns,
	).FROM(User).WHERE(
		User.ID.EQ(Int32(int32(id))),
	)

	err := stmt.Query(db, &user)

	return user, err
}

func (u UserRepository) FetchByUsername(username string) (model.User, error) {
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

func (u UserRepository) FetchSecurityStamp(userid int) (string, error) {
	stamp := ""

	stmt := SELECT(
		User.SecurityStamp,
	).FROM(User).WHERE(
		User.ID.EQ(Int32(int32(userid))),
	)

	err := stmt.Query(db, &stamp)
	return stamp, err
}

func (u UserRepository) Update(user model.User) (sql.Result, error) {
	stmt := User.UPDATE(User.MutableColumns).
		MODEL(user).
		WHERE(User.ID.EQ(Int32(user.ID)))

	return stmt.Exec(db)
}
