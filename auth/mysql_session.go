package auth

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gohttp/config"
	"log"
	"net/http"
	"time"
)

// Implements a SessionStore interface
type MySqlSessionStore struct {
	base *sessionStoreBase
	db   *sql.DB
}

func (st *MySqlSessionStore) InitStore(name string,
	itemExpiry time.Duration,
	willRedirect bool,
	loginPath string,
	logoutPath string,
	defaultPath string) {
	st.base = &sessionStoreBase{}
	st.base.name = name
	st.base.ctxKey = sessionKey{}
	st.base.expiration = itemExpiry
	st.base.willRedirect = willRedirect
	st.base.LoginPath = loginPath
	st.base.LogoutPath = logoutPath
	st.base.DefaultPath = defaultPath

	config := config.GetConfiguration()

	var err error

	st.db, err = sql.Open("mysql", config.ConnectionString)

	if err != nil {
		panic(err.Error())
	}

	log.Printf("Initialized MySQL session authentication [redirects: %t]\n", willRedirect)
}

func (st *MySqlSessionStore) LoadSession(next http.Handler, requireAuth bool) http.Handler {
	var id *Identity

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id = st.GetIdentityFromRequest(w, r)

		handler := st.base.loadSession(next, id, requireAuth)
		handler.ServeHTTP(w, r)
	})
}

func (st *MySqlSessionStore) PutSession(w http.ResponseWriter, r *http.Request, id *Identity) {
	cookieValue := randBase64String(cookieEntropy)

	sql := "INSERT INTO sessions (id, userid, useragent, ipaddr, logintime) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := st.db.Prepare(sql)

	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(cookieValue, id.UserId, id.UserAgent, id.IpAddr, id.LoginTime)

	if err != nil {
		log.Println(err)
	}

	st.base.setCookie(w, r, cookieValue, id.RememberMe)
}

func (st *MySqlSessionStore) DeleteSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(st.base.name)
	if err != nil {
		return
	}

	sql := "DELETE FROM sessions WHERE id = ?"
	stmt, _ := st.db.Prepare(sql)

	stmt.Exec(cookie.Value)

	st.base.removeCookie(w, r)
}

func (st *MySqlSessionStore) GetIdentityFromRequest(w http.ResponseWriter, r *http.Request) *Identity {
	cookie, err := r.Cookie(st.base.name)

	if err != nil {
		return nil
	}

	var id Identity

	sql := "SELECT userid, useragent, ipaddr, logintime FROM sessions WHERE id = ?"

	stmt, err := st.db.Prepare(sql)

	if err != nil {
		log.Println(err)
	}

	err1 := stmt.QueryRow(cookie.Value).Scan(&id.UserId, &id.UserAgent, &id.IpAddr, &id.LoginTime)

	if err1 != nil {
		log.Println(err1)
		return nil
	}

	id.IsAuthenticated = true //must manually set true since database does not store the entire Identity struct

	return &id
}

func (st *MySqlSessionStore) GetIdentityFromCtx(r *http.Request) *Identity {
	return st.base.getIdentityFromCtx(r)
}

func (st *MySqlSessionStore) GetAllIdentities(id *Identity) []Identity {
	identities := make([]Identity, 0)

	sql := "SELECT userid, useragent, ipaddr, logintime FROM sessions WHERE userid = ?"

	stmt, err := st.db.Prepare(sql)

	if err != nil {
		log.Println(err)
	}

	rows, err := stmt.Query(id.UserId)
	defer rows.Close()

	for rows.Next() {
		var id Identity
		rows.Scan(&id.UserId, &id.UserAgent, &id.IpAddr, &id.LoginTime)

		identities = append(identities, id)
	}

	return identities
}

func (st *MySqlSessionStore) GetBase() *sessionStoreBase {
	return st.base
}

func (st *MySqlSessionStore) DeleteAllByUserId(w http.ResponseWriter, r *http.Request, id *Identity) {
	sql := "DELETE FROM sessions WHERE userid = ?"

	stmt, err := st.db.Prepare(sql)

	if err != nil {
		log.Println(err)
	}

	_, err1 := stmt.Exec(id.UserId)

	if err1 != nil {
		log.Println(err1)
	}

	st.base.removeCookie(w, r)
}
