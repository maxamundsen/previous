package auth

import (
	"webdawgengine/database"
	"webdawgengine/identity"
	"log"
	"net/http"
	"time"
)

// Implements a SessionStore interface
type MySqlSessionStore struct {
	base *sessionStoreBase
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

	log.Printf("Initialized MySQL session authentication [redirects: %t]\n", willRedirect)
}

func (st *MySqlSessionStore) LoadSession(h http.HandlerFunc, requireAuth bool) http.HandlerFunc {
	var id *identity.Identity

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id = st.GetIdentityFromRequest(w, r)

		handler := st.base.loadSession(h, id, requireAuth)
		handler.ServeHTTP(w, r)
	})
}

func (st *MySqlSessionStore) PutSession(w http.ResponseWriter, r *http.Request, id *identity.Identity) {
	id.Id = randBase64String(cookieEntropy)

	err := database.InsertIdentity(id)
	if err != nil {
		log.Println(err)
	}

	time.AfterFunc(st.base.expiration, func() {
		st.DeleteSessionByKey(id.Id)
	})

	st.base.setCookie(w, r, id.Id, id.RememberMe)
}

func (st *MySqlSessionStore) DeleteSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(st.base.name)
	if err != nil {
		return
	}

	st.DeleteSessionByKey(cookie.Value)

	st.base.removeCookie(w, r)
}

func (st *MySqlSessionStore) DeleteSessionByKey(sessionKey string) {
	err := database.DeleteIdentityById(sessionKey)
	if err != nil {
		log.Println(err)
	}
}

func (st *MySqlSessionStore) GetIdentityFromRequest(w http.ResponseWriter, r *http.Request) *identity.Identity {
	cookie, err := r.Cookie(st.base.name)
	if err != nil {
		return nil
	}

	id, fetchErr := database.FetchIdentityById(cookie.Value)
	if fetchErr != nil {
		log.Println(fetchErr)
		return nil
	}

	// fetch claims each request
	claims, claimsErr := database.FetchUserClaimsByEmail(id.Email)

	if claimsErr != nil {
		log.Println(claimsErr)
		return nil
	}

	id.Claims = claims

	id.IsAuthenticated = true //must manually set true since database does not store the entire Identity struct

	return &id
}

func (st *MySqlSessionStore) GetIdentityFromCtx(r *http.Request) *identity.Identity {
	return st.base.getIdentityFromCtx(r)
}

func (st *MySqlSessionStore) GetAllIdentities(id *identity.Identity) []identity.Identity {
	identities, err := database.FetchAllIdentitiesByEmail(id.Email)
	if err != nil {
		log.Println(err)
	}

	return identities
}

func (st *MySqlSessionStore) GetBase() *sessionStoreBase {
	return st.base
}

func (st *MySqlSessionStore) DeleteAllByEmail(w http.ResponseWriter, r *http.Request, id *identity.Identity) {
	err := database.DeleteAllIdentitiesByEmail(id.Email)
	if err != nil {
		log.Println(err)
	}

	st.base.removeCookie(w, r)
}
