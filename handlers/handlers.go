package handlers

import (
	"gohttp/auth"
	"gohttp/config"
	"time"
)

var sessionStore auth.SessionStore

// Initialize a session implementation and put it in the package scope
// so it can be used by handler functions
func SessionInit() {
	config := config.GetConfiguration()

	sessionStore = &auth.MemorySessionStore{}
	cookieExpiry := time.Duration(time.Hour * 24 * time.Duration(config.CookieExpiryDays))

	// set store options here
	// since the routes are determined at compile time, not runtime, it would not make
	// sense to put these options in the config file
	sessionStore.InitStore("id", cookieExpiry, true, "/auth/login", "/auth/logout", "/example")
}
