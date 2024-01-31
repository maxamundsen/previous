package handlers

import (
	"log"
	"net/http"
)

// Dynamic routes are defined here
func MapDynamicRoutes(mux *http.ServeMux) {
	// root (index) handler
	mux.HandleFunc("/", indexHandler)

	// example handlers
	mux.Handle("/example", handleSession(exampleHandler, true))
	mux.Handle("/example/upload", handleSession(exampleUploadHandler, true))
	mux.Handle("/example/counter", handleSession(exampleCounterHandler, true))
	mux.Handle("/example/passgen", handleSession(examplePassgenHandler, true))
	mux.Handle("/example/database", handleSession(exampleDatabaseHandler, true))
	mux.Handle("/example/adduser", handleSession(exampleAdduserHandler, true))
	mux.Handle("/example/deleteall", handleSession(exampleDeleteallHandler, true))
	mux.Handle("/example/mail", handleSession(exampleMailHandler, true))
	mux.Handle("/example/fetch", handleSession(exampleFetchHandler, true))

	// auth handlers
	mux.Handle("/auth/login", handleSession(loginHandler, false))
	mux.Handle("/auth/logout", handleSession(logoutHandler, true))
	mux.Handle("/auth/logoutall", handleSession(logoutAllHandler, true))

	// account handlers
	mux.Handle("/account/sessions", handleSession(accountSessionHandler, true))
	mux.Handle("/account/info", handleSession(accountInfoHandler, true))

	// api handlers
	mux.Handle("/api/identity", handleSession(apiIdentityHandler, true))

	log.Println("Mapped dynamic routes")
}

// middleware wrappers
func handleSession(handlerFunc http.HandlerFunc, isAuth bool) http.Handler {
	return sessionStore.LoadSession(http.HandlerFunc(handlerFunc), isAuth)
}
