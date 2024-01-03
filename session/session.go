package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
)

type AuthSession struct {
	IsAuthenticated bool
	Role     string
	Username string
}

type sessionKey struct{}

// SessionStore holds the session data and settings
type SessionStore struct {
	name          string
	sessions      map[string]*AuthSession
	lock          sync.RWMutex
	ctxKey        sessionKey
	expiration    time.Duration
	willRedirect  bool
	redirectPath  string
	defaultPath   string
}

// Init will initialize the SessionStore object
func (st *SessionStore) InitStore(name string, itemExpiry time.Duration, willRedirect bool, redirectPath string, defaultPath string) {
	st.name = name
	st.sessions = make(map[string]*AuthSession)
	st.ctxKey = sessionKey{}
	st.expiration = itemExpiry
	st.willRedirect = willRedirect
	st.redirectPath = redirectPath
	st.defaultPath = defaultPath
}

func randBase64String(entropyBytes int) string {
	b := make([]byte, entropyBytes)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

// PutSession will store the session in the SessionStore.
// The session will automatically expire after defined SessionStore.sessionExpiration.
func (st *SessionStore) PutSession(w http.ResponseWriter, r *http.Request, sess *AuthSession) {
	cookieValue := randBase64String(33) // 33 bytes entropy

	time.AfterFunc(st.expiration, func() {
		st.lock.Lock()
		delete(st.sessions, cookieValue)
		st.lock.Unlock()
	})

	st.lock.Lock()
	st.sessions[cookieValue] = sess
	st.lock.Unlock()

	cookie := &http.Cookie{
		Name:     st.name,
		Value:    cookieValue,
		Expires:  time.Now().Add(st.expiration),
		HttpOnly: true,
		Secure:   r.URL.Scheme == "https",
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}
	
	http.SetCookie(w, cookie)
}

// DeleteSession will delete the session from the SessionStore.
func (st *SessionStore) DeleteSession(r *http.Request) {
	cookie, err := r.Cookie(st.name)
	if err != nil {
		return
	}
	st.lock.Lock()
	delete(st.sessions, cookie.Value)
	st.lock.Unlock()
}

// GetSessionFromRequest retrieves the session from the http.Request cookies.
// The function will return nil if the session does not exist within the http.Request cookies.
func (st *SessionStore) GetSessionFromRequest(r *http.Request) *AuthSession {
	cookie, err := r.Cookie(st.name)
	if err != nil {
		return nil
	}
	st.lock.RLock()
	sess := st.sessions[cookie.Value]
	st.lock.RUnlock()
	return sess
}

// LoadSession will load the session into the http.Request context.
func (st *SessionStore) LoadSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess := st.GetSessionFromRequest(r)
		
		if sess == nil  {
			noAuthSession := &AuthSession {
				IsAuthenticated: false,
			}
			
			ctx := context.WithValue(r.Context(), st.ctxKey, noAuthSession)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		
		ctx := context.WithValue(r.Context(), st.ctxKey, sess)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetSessionFromCtx retrieves the session from the http.Request context.
// If no session is found, it returns an AuthSession with IsAuthenticated set to false.
func (st *SessionStore) GetSessionFromCtx(r *http.Request) *AuthSession {
	return r.Context().Value(st.ctxKey).(*AuthSession)
}

// Check session for auth, handle accordingly
func (st *SessionStore) AuthorizeRoute(w http.ResponseWriter, r *http.Request, a *AuthSession) {
	if !a.IsAuthenticated {
		if st.willRedirect && st.redirectPath != r.URL.Path {
			http.Redirect(w, r, st.redirectPath, http.StatusFound)
		} else if !st.willRedirect {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	} else {
		if st.willRedirect && st.redirectPath == r.URL.Path {
			http.Redirect(w, r, st.defaultPath, http.StatusFound)
		}
	}
}