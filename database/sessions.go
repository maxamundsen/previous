package database

import (
	"webdawgengine/identity"
)

func FetchIdentityById(sessionId string) (identity.Identity, error) {
	var id identity.Identity

	sql := "SELECT id, email, useragent, ipaddr, logintime FROM sessions WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return id, err
	}

	err1 := stmt.QueryRow(sessionId).Scan(&id.Id, &id.Email, &id.UserAgent, &id.IpAddr, &id.LoginTime)
	if err1 != nil {
		return id, err1
	}

	return id, nil
}

func DeleteAllIdentitiesByEmail(email string) error {
	sql := "DELETE FROM sessions WHERE email = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}

	_, err1 := stmt.Exec(email)

	if err1 != nil {
		return err1
	}

	return nil
}

func FetchAllIdentitiesByEmail(email string) ([]identity.Identity, error) {
	identities := make([]identity.Identity, 0)

	sql := "SELECT id, email, useragent, ipaddr, logintime FROM sessions WHERE email = ? ORDER BY logintime DESC"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return identities, err
	}

	rows, queryErr := stmt.Query(email)
	if queryErr != nil {
		return identities, queryErr
	}

	defer rows.Close()

	for rows.Next() {
		var identity identity.Identity
		rows.Scan(&identity.Id, &identity.Email, &identity.UserAgent, &identity.IpAddr, &identity.LoginTime)
		identities = append(identities, identity)
	}

	return identities, nil
}

func DeleteIdentityById(sessionId string) error {
	sql := "DELETE FROM sessions WHERE id = ?"

	stmt, err:= db.Prepare(sql)
	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(sessionId)
	if execErr != nil {
		return execErr
	}

	return nil
}

func InsertIdentity(id *identity.Identity) error{
	sql := "INSERT INTO sessions (id, email, useragent, ipaddr, logintime) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}

	_, execErr := stmt.Exec(id.Id, id.Email, id.UserAgent, id.IpAddr, id.LoginTime)

	if execErr != nil {
		return execErr
	}

	return nil
}