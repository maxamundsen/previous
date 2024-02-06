package handlers

import (
	"webdawgengine/auth"
	"webdawgengine/config"
	"time"
)

// the handlers package contains definitions for each HTTP endpoint,
// and which function gets called when that endpoint is hit. middleware
// can be setup to intercept the requests, do some task, then pass the request down
// to the individual endpoint handler. for example, when an route is wrapped with
// session middleware, it checks for a valid session *every request* based on a
// value passed via a cookie in the request. middleware allows us to automatically handle
// incoming requests without having to call into any extra functions within the handler itself.

// a session store struct is scoped to the handlers package to be used directly
// within the handler functions.

var sessionStore auth.SessionStore

func SessionInit() {
	// sessionStore = &auth.MemorySessionStore{}
	sessionStore = &auth.MySqlSessionStore{}
	cookieExpiry := time.Duration(time.Hour * 24 * time.Duration(config.GetCookieExpiryDays()))
	sessionStore.InitStore("id", cookieExpiry, true, "/auth/login", "/auth/logout", "/example")
}
