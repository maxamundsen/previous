package auth

import (
	"log"
	"net/http"
	"sync"
	"time"
	"webdawgengine/identity"
)

// Implements a SessionStore interface
type MemorySessionStore struct {
	base     *sessionStoreBase
	sessions map[string]*identity.Identity
	lock     sync.RWMutex
}

func (st *MemorySessionStore) InitStore(name string,
	itemExpiry time.Duration,
	willRedirect bool,
	loginPath string,
	logoutPath string,
	defaultPath string) {
	st.base = &sessionStoreBase{}
	st.sessions = make(map[string]*identity.Identity)
	st.base.name = name
	st.base.ctxKey = sessionKey{}
	st.base.expiration = itemExpiry
	st.base.willRedirect = willRedirect
	st.base.LoginPath = loginPath
	st.base.LogoutPath = logoutPath
	st.base.DefaultPath = defaultPath
	log.Printf("Initialized in-memory session authentication [redirects: %t]\n", willRedirect)
}

func (st *MemorySessionStore) LoadSession(h http.HandlerFunc, requireAuth bool) http.HandlerFunc {
	var id *identity.Identity

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id = st.GetIdentityFromRequest(w, r)

		handler := st.base.loadSession(h, id, requireAuth)
		handler.ServeHTTP(w, r)
	})
}

func (st *MemorySessionStore) PutSession(w http.ResponseWriter, r *http.Request, id *identity.Identity) {
	cookieValue := randBase64String(cookieEntropy) // sets the entropy of the random string

	// Delete the session from the store after expiration time
	time.AfterFunc(st.base.expiration, func() {
		st.DeleteSessionByKey(cookieValue)
	})

	st.lock.Lock()
	st.sessions[cookieValue] = id
	st.lock.Unlock()

	st.base.setCookie(w, r, cookieValue, id.RememberMe)
}

func (st *MemorySessionStore) DeleteSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(st.base.name)
	if err != nil {
		return
	}

	st.DeleteSessionByKey(cookie.Value)

	st.base.removeCookie(w, r)
}

func (st *MemorySessionStore) DeleteSessionByKey(sessionKey string) {
	st.lock.Lock()
	delete(st.sessions, sessionKey)
	st.lock.Unlock()
}

func (st *MemorySessionStore) GetIdentityFromRequest(w http.ResponseWriter, r *http.Request) *identity.Identity {
	cookie, err := r.Cookie(st.base.name)

	if err != nil {
		return nil
	}

	st.lock.RLock()
	id := st.sessions[cookie.Value]
	st.lock.RUnlock()

	return id
}

func (st *MemorySessionStore) GetIdentityFromCtx(r *http.Request) *identity.Identity {
	return st.base.getIdentityFromCtx(r)
}

func (st *MemorySessionStore) GetAllIdentities(id *identity.Identity) []identity.Identity {
	identities := make([]identity.Identity, 0)

	// find all sessions that contain the input identity
	for _, v := range st.sessions {
		if v.Email == id.Email {
			identities = append(identities, *v)
		}
	}

	return identities
}

func (st *MemorySessionStore) GetBase() *sessionStoreBase {
	return st.base
}

func (st *MemorySessionStore) DeleteAllByEmail(w http.ResponseWriter, r *http.Request, id *identity.Identity) {
	for i, v := range st.sessions {
		if v.Email == id.Email {
			st.lock.Lock()
			delete(st.sessions, i)
			st.lock.Unlock()
		}
	}

	st.base.removeCookie(w, r)
}
