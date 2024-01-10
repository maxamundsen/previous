package auth

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// Preface:
// This file contains the basis for creating "authentication sessions".
// Other web frameworks contain crazy complecated authentication middleware and
// identity management. This one should be pretty simple to understand.
// Auth session management does NOT contain any code to actually authenticate users
// (username/password checking, password hashing etc)

// 1. Users
// In this simple authentication system, there is an Identity struct that represents
// a "user" on the system. You can customize this structure to fit your needs.

// 2. Stores
// SessionStore is an interface that describes the capabilities of a 'session store'
// A session store is a storage mechanism for storing Identity structures.
// The type of storage implemented does not matter, as long as it can 


type SessionStore interface {
	InitStore(name string, 
	          itemExpiry time.Duration, 
	          willRedirect bool, 
	          loginPath string, 
	          logoutPath string, 
	          defaultPath string)
	PutSession(w http.ResponseWriter, r *http.Request, id *Identity)
	DeleteSession(w http.ResponseWriter, r *http.Request)
	LoadSession(next http.Handler, requireAuth bool) http.Handler
	GetSessionFromCtx(r *http.Request) *Identity
	GetSessionFromRequest(r *http.Request) *Identity
	GetBase() *sessionStoreBase
}


type sessionStoreBase struct {
	name         string
	ctxKey       sessionKey
	expiration   time.Duration
	WillRedirect bool
	LoginPath    string
	LogoutPath   string
	DefaultPath  string
}

type sessionKey struct{}

func (st *sessionStoreBase) setCookie(w http.ResponseWriter,
                                      r *http.Request, 
                                      cookieValue string, 
                                      rememberMe bool) {
	cookie := &http.Cookie{
		Name:     st.name,
		Value:    cookieValue,
		HttpOnly: true,
		Secure:   r.URL.Scheme == "https",
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	// if no expiry is set, cookie defaults to clear after browser closes
	if rememberMe {
		cookie.Expires = time.Now().Add(st.expiration)
	}

	http.SetCookie(w, cookie)
}

func (st *sessionStoreBase) removeCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
    	Name: st.name,
    	MaxAge: -1,
    	Expires: time.Now().Add(-100 * time.Hour),// Set expires for older versions of IE
    	Path: "/",
	})
}

// middleware for loading a provided auth session, and automatically
// handling redirections
func (st *sessionStoreBase) loadSession(next http.Handler, 
                                        id *Identity, 
                                        requireAuth bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if not auth'd
		if id == nil {
			blankIdentity := &Identity{IsAuthenticated: false}
	
			if requireAuth {
				if st.WillRedirect && st.LoginPath != r.URL.Path && st.LogoutPath != r.URL.Path {
					redirectPath := st.LoginPath + "?redirect=" + url.QueryEscape(r.URL.String())
	
					http.Redirect(w, r, redirectPath, http.StatusFound)
					return
				} else if !st.WillRedirect && st.LoginPath != r.URL.Path {
					http.Error(w, "Error: Unauthorized", http.StatusUnauthorized)
					return
				}
			}
	
			ctx := context.WithValue(r.Context(), st.ctxKey, blankIdentity)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	
		// if auth'd
		if st.WillRedirect && st.LoginPath == r.URL.Path {
			http.Redirect(w, r, st.DefaultPath, http.StatusFound)
			return
		}
	
		// if there is a valid identity
		ctx := context.WithValue(r.Context(), st.ctxKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))	
	})
}

// returns a session found in the http request context
func (st *sessionStoreBase) getSessionFromCtx(r *http.Request) *Identity {
	return r.Context().Value(st.ctxKey).(*Identity)
}
