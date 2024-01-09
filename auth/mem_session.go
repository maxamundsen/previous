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
	"log"
	"net/http"
	"sync"
	"time"
)

// Implements a SessionStore interface
type MemorySessionStore struct {
	base     *sessionStoreBase
	sessions map[string]*AuthSession
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
	st.sessions = make(map[string]*AuthSession)
	st.base.name = name
	st.base.ctxKey = sessionKey{}
	st.base.expiration = itemExpiry
	st.base.WillRedirect = willRedirect
	st.base.LoginPath = loginPath
	st.base.LogoutPath = logoutPath
	st.base.DefaultPath = defaultPath
	log.Printf("Initialized in-memory session authentication [redirects: %t]\n", willRedirect)
}

// middleware for loading sessions
func (st *MemorySessionStore) LoadSession(next http.Handler, requireAuth bool) http.Handler {
	var sess *AuthSession
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess = st.GetSessionFromRequest(r)
		
		handler := st.base.loadSession(next, sess, requireAuth)	
		handler.ServeHTTP(w, r)
	})
}

// PutSession will store the session in the MemorySessionStore.
// The session will automatically expire after defined MemorySessionStore.sessionExpiration.
func (st *MemorySessionStore) PutSession(w http.ResponseWriter, r *http.Request, sess *AuthSession) {
	cookieValue := randBase64String(constants.CookieEntropy)

	// Delete the session from the store after expiration time
	time.AfterFunc(st.base.expiration, func() {
		st.lock.Lock()
		delete(st.sessions, cookieValue)
		st.lock.Unlock()
	})

	st.lock.Lock()
	st.sessions[cookieValue] = sess
	st.lock.Unlock()

	st.base.setCookie(w, r, cookieValue, sess.RememberMe)
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

// getSessionFromRequest retrieves the session from the http.Request cookies.
// The function will return nil if the session does not exist within the http.Request cookies.
func (st *MemorySessionStore) GetSessionFromRequest(r *http.Request) *AuthSession {
	cookie, err := r.Cookie(st.base.name)
	if err != nil {
		return nil
	}
	st.lock.RLock()
	sess := st.sessions[cookie.Value]
	st.lock.RUnlock()
	return sess
}

func (st *MemorySessionStore) GetSessionFromCtx(r *http.Request) *AuthSession {
	return st.base.getSessionFromCtx(r)
}

func (st *MemorySessionStore) GetBase() *sessionStoreBase {
	return st.base
}