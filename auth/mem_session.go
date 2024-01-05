package auth

// mem_session provides functions for handling an in-memory session store.
// keys are retrieved from a browser cookie to obtain a session struct
// from a map in memory. sessions can be created or destroyed from http handlers
// (such as a login page)

// since the session store (which contains the session map)  is a struct in memory,
// it is destroyed  when the program exits. for longer lived sessions, a database
// store is preferred.

import (
	"gohttp/constants"
	"net/http"
	"sync"
	"log"
	"time"
)

type MemorySessionStore struct {
	Base     *SessionStore
	sessions map[string]*AuthSession
	lock     sync.RWMutex
}

// Init will initialize the MemorySessionStore object
func (st *MemorySessionStore) InitStore(name string, itemExpiry time.Duration, WillRedirect bool, LoginPath string, LogoutPath string, DefaultPath string) {
	st.Base = &SessionStore{}
	st.sessions = make(map[string]*AuthSession)
	st.Base.name = name
	st.Base.ctxKey = SessionKey{}
	st.Base.expiration = itemExpiry
	st.Base.WillRedirect = WillRedirect
	st.Base.LoginPath = LoginPath
	st.Base.LogoutPath = LogoutPath
	st.Base.DefaultPath = DefaultPath
	log.Printf("Initialized in-memory session authentication [redirects: %t]\n", WillRedirect)
}

// PutSession will store the session in the MemorySessionStore.
// The session will automatically expire after defined MemorySessionStore.sessionExpiration.
func (st *MemorySessionStore) PutSession(w http.ResponseWriter, r *http.Request, sess *AuthSession) {
	cookieValue := randBase64String(constants.CookieEntropy)

	// Delete the session from the store after expiration time
	time.AfterFunc(st.Base.expiration, func() {
		st.lock.Lock()
		delete(st.sessions, cookieValue)
		st.lock.Unlock()
	})

	st.lock.Lock()
	st.sessions[cookieValue] = sess
	st.lock.Unlock()

	st.Base.setCookie(w, r, cookieValue, sess.RememberMe)
}

// DeleteSession will delete the session from the MemorySessionStore.
func (st *MemorySessionStore) DeleteSession(r *http.Request) {
	cookie, err := r.Cookie(st.Base.name)
	if err != nil {
		return
	}
	st.lock.Lock()
	delete(st.sessions, cookie.Value)
	st.lock.Unlock()
}

func (st *MemorySessionStore) LoadSession(next http.Handler, requireAuth bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := st.GetSessionFromRequest(r)
		st.Base.loadSession(next, w, r, sess, requireAuth)
	})
}

// GetSessionFromRequest retrieves the session from the http.Request cookies.
// The function will return nil if the session does not exist within the http.Request cookies.
func (st *MemorySessionStore) GetSessionFromRequest(r *http.Request) *AuthSession {
	cookie, err := r.Cookie(st.Base.name)
	if err != nil {
		return nil
	}
	st.lock.RLock()
	sess := st.sessions[cookie.Value]
	st.lock.RUnlock()
	return sess
}
