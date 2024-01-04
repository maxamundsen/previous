package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"time"
)

type SessionStore struct {
	name         string
	ctxKey       SessionKey
	expiration   time.Duration
	WillRedirect bool
	LoginPath    string
	LogoutPath   string
	DefaultPath  string
}

type AuthSession struct {
	IsAuthenticated bool
	RememberMe      bool
	Role            string
	Username        string
}

type SessionKey struct{}

func (st *SessionStore) setCookie(w http.ResponseWriter, r *http.Request, cookieValue string, rememberMe bool) {
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

func (st *SessionStore) loadSession(next http.Handler, w http.ResponseWriter, r *http.Request, sess *AuthSession, requireAuth bool) {
	// if not auth'd
	if sess == nil {
		noAuthSession := &AuthSession{IsAuthenticated: false}

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

		ctx := context.WithValue(r.Context(), st.ctxKey, noAuthSession)
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	}
	
	// if auth'd
	if st.WillRedirect && st.LoginPath == r.URL.Path {
		http.Redirect(w, r, st.DefaultPath, http.StatusFound)
		return
	}
	
	// if there is a valid session
	ctx := context.WithValue(r.Context(), st.ctxKey, sess)
	next.ServeHTTP(w, r.WithContext(ctx))
}

// GetSessionFromCtx retrieves the session from the http.Request context.
// If no session is found, it returns an AuthSession with IsAuthenticated set to false.
func (st *SessionStore) GetSessionFromCtx(r *http.Request) *AuthSession {
	return r.Context().Value(st.ctxKey).(*AuthSession)
}

func randBase64String(entropyBytes int) string {
	b := make([]byte, entropyBytes)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
