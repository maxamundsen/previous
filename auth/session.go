package auth

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

// This file contains the basis for creating "authentication sessions".
// Other web frameworks contain crazy complecated authentication middleware and
// identity management. This one should be pretty simple to understand.
// Auth session management does NOT contain any code to actually authenticate users
// (username/password checking, password hashing etc). That should be handled elsewhere.

// In this simple authentication system, there is an Identity struct that represents
// a "user" on the system. You can customize this structure to fit your needs.

// Identities are not stored directly in the browser cookies.
// Instead, they are stored on the server in a key-value pair called a 'session store'.
// When a user successfully authenticates, an entry is made in the store
// containing an Identity, and a randomly generated base64 string key. The
// key is appended to the response as a cookie, and stored in the users browser.

// SessionStore is an interface that describes the capabilities of a session store.
// The type of storage implemented does not matter, as long as the custom storage
// type contains all of the methods in the interface, and a sessionStoreBase.

// Whatever implementation of the interface you choose to use, the method signatures will
// always be the same. Http endpoint handlers that wish to use sessions do so without knowing
// the implementation.

type SessionStore interface {
	InitStore(name string,
		itemExpiry time.Duration,
		willRedirect bool,
		loginPath string,
		logoutPath string,
		defaultPath string)
	PutSession(w http.ResponseWriter, r *http.Request, id *Identity)
	DeleteSession(w http.ResponseWriter, r *http.Request)
	DeleteAllByUserId(w http.ResponseWriter, r *http.Request, id *Identity)
	LoadSession(next http.Handler, requireAuth bool) http.Handler
	GetIdentityFromCtx(r *http.Request) *Identity
	GetIdentityFromRequest(w http.ResponseWriter, r *http.Request) *Identity
	GetAllIdentities(id *Identity) []Identity
	GetBase() *sessionStoreBase
}

// The base store struct contains basic properties of a session store.
type sessionStoreBase struct {
	name         string
	ctxKey       sessionKey
	expiration   time.Duration
	willRedirect bool   // used to determine if unauthorized requests get a 401, or redirect
	LoginPath    string // redirect path if unauthorized
	LogoutPath   string
	DefaultPath  string // redirect path if authorized
}

type sessionKey struct{}

// hardcoded value of 33, feel free to modify
const cookieEntropy int = 33

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
		// expiration is set from the config file
		cookie.Expires = time.Now().Add(st.expiration)
	}

	http.SetCookie(w, cookie)
}

// create a cookie with the same name, but with no value, then append it to the response
// setting a 'blank' cookie will delete it from the browser
func (st *sessionStoreBase) removeCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    st.name,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour), // Set expires for older versions of IE
		Path:    "/",
	})
}

// middleware for loading a provided auth session, and automatically
// handling redirections
func (st *sessionStoreBase) loadSession(next http.Handler,
	id *Identity,
	requireAuth bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if id == nil {
			blankIdentity := &Identity{IsAuthenticated: false}

			if requireAuth {
				if st.willRedirect && st.LoginPath != r.URL.Path && st.LogoutPath != r.URL.Path {
					redirectPath := st.LoginPath + "?redirect=" + url.QueryEscape(r.URL.String())

					http.Redirect(w, r, redirectPath, http.StatusFound)
					return
				} else if !st.willRedirect && st.LoginPath != r.URL.Path {
					http.Error(w, "Error: Unauthorized", http.StatusUnauthorized)
					return
				}
			}

			ctx := context.WithValue(r.Context(), st.ctxKey, blankIdentity)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if st.willRedirect && st.LoginPath == r.URL.Path {
			http.Redirect(w, r, st.DefaultPath, http.StatusFound)
			return
		}

		ctx := context.WithValue(r.Context(), st.ctxKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// returns an identity found in the http request context
func (st *sessionStoreBase) getIdentityFromCtx(r *http.Request) *Identity {
	return r.Context().Value(st.ctxKey).(*Identity)
}
