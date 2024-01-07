package handlers

import (
	"gohttp/auth"
	"gohttp/constants"
)

var MemorySession auth.MemorySessionStore

func SessionInit() {
	MemorySession.InitStore("AuthenticationCookie", constants.CookieExpiryTime, true, "/login", "/logout", "/test")
}
