package handlers

import (
	"gohttp/auth"
	"gohttp/config"
	"time"
)

var sessionStore auth.SessionStore

func SessionInit() {
	config := config.GetConfiguration()

	sessionStore = &auth.MemorySessionStore{}
	cookieExpiry := time.Duration(time.Hour * 24 * time.Duration(config.CookieExpiryDays))
	sessionStore.InitStore("ID", cookieExpiry, true, "/login", "/logout", "/test")
}
