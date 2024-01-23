package handlers

import (
	"gohttp/auth"
	"time"
)

var sessionStore auth.SessionStore

// Initialize a session implementation and put it in the package scope
// so it can be used by handler functions
func SessionInit(cookieExpiryDays int) {
	// sessionStore = &auth.MemorySessionStore{}
	sessionStore = &auth.MySqlSessionStore{}
	cookieExpiry := time.Duration(time.Hour * 24 * time.Duration(cookieExpiryDays))
	sessionStore.InitStore("id", cookieExpiry, true, "/auth/login", "/auth/logout", "/example")
}
