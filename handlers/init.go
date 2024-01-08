package handlers

import (
	"gohttp/auth"
	"gohttp/constants"
)

var sessionStore auth.SessionStore

func SessionInit() {
	sessionStore = &auth.MemorySessionStore{}
	sessionStore.InitStore("ID", constants.CookieExpiryTime, true, "/login", "/logout", "/test")
}
