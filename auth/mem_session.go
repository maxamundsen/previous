package auth

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// Implements a SessionStore interface
type MemorySessionStore struct {
	base     *sessionStoreBase
	sessions map[string]*Identity
	lock     sync.RWMutex
}

// Init will initialize the MemorySessionStore object
func (st *MemorySessionStore) InitStore(name string,
	itemExpiry time.Duration,
	willRedirect bool,
	loginPath string,
	logoutPath string,
	defaultPath string) {
	st.base = &sessionStoreBase{}
	st.sessions = make(map[string]*Identity)
	st.base.name = name
	st.base.ctxKey = sessionKey{}
	st.base.expiration = itemExpiry
	st.base.willRedirect = willRedirect
	st.base.LoginPath = loginPath
	st.base.LogoutPath = logoutPath
	st.base.DefaultPath = defaultPath
	log.Printf("Initialized in-memory session authentication [redirects: %t]\n", willRedirect)
}

// middleware for loading sessions
func (st *MemorySessionStore) LoadSession(next http.Handler, requireAuth bool) http.Handler {
	var id *Identity

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id = st.GetIdentityFromRequest(w, r)

		handler := st.base.loadSession(next, id, requireAuth)
		handler.ServeHTTP(w, r)
	})
}

// PutSession will store the session in the MemorySessionStore.
// The session will automatically expire after defined MemorySessionStore.sessionExpiration.
func (st *MemorySessionStore) PutSession(w http.ResponseWriter, r *http.Request, id *Identity) {
	cookieValue := randBase64String(cookieEntropy) // sets the entropy of the random string

	// Delete the session from the store after expiration time
	time.AfterFunc(st.base.expiration, func() {
		st.lock.Lock()
		delete(st.sessions, cookieValue)
		st.lock.Unlock()
	})

	st.lock.Lock()
	st.sessions[cookieValue] = id
	st.lock.Unlock()

	st.base.setCookie(w, r, cookieValue, id.RememberMe)
}

// DeleteSession will delete the session from the MemorySessionStore.
func (st *MemorySessionStore) DeleteSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(st.base.name)
	if err != nil {
		return
	}
	st.lock.Lock()
	delete(st.sessions, cookie.Value)
	st.lock.Unlock()

	st.base.removeCookie(w, r)
}

// getIdentityFromRequest retrieves the Identity from the http.Request cookies.
// The function will return nil if the session does not exist within the http.Request cookies.
func (st *MemorySessionStore) GetIdentityFromRequest(w http.ResponseWriter, r *http.Request) *Identity {
	cookie, err := r.Cookie(st.base.name)
	if err != nil {
		return nil
	}

	st.lock.RLock()
	id := st.sessions[cookie.Value]
	st.lock.RUnlock()

	return id
}

func (st *MemorySessionStore) GetIdentityFromCtx(r *http.Request) *Identity {
	return st.base.getIdentityFromCtx(r)
}

func (st *MemorySessionStore) GetAllIdentities(id *Identity) []*Identity {
	identities := make([]*Identity, 0)

	// find all sessions that contain the input identity
	for _, v := range st.sessions {
		if v.UserId == id.UserId {
			identities = append(identities, v)
		}
	}

	return identities
}

func (st *MemorySessionStore) GetBase() *sessionStoreBase {
	return st.base
}

// delete all other auth sessions containing the provided identity
// make to to log user out when completed for security reasons
// an attacker should not be able to log out 'other' sessions without also
// needing to log back in
func (st *MemorySessionStore) DeleteAllById(w http.ResponseWriter, r *http.Request, id *Identity) {
	for i, v := range st.sessions {
		if v.UserId == id.UserId {
			st.lock.Lock()
			delete(st.sessions, i)
			st.lock.Unlock()
		}
	}

	st.base.removeCookie(w, r)
}
