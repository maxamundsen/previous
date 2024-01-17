package handlers

import (
	"log"
	"net/http"
)

// Dynamic routes are defined here
func MapDynamicRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)

	mux.Handle("/example", sessionStore.LoadSession(http.HandlerFunc(exampleHandler), true))
	mux.Handle("/example/counter", sessionStore.LoadSession(http.HandlerFunc(exampleCounterHandler), true))
	mux.Handle("/example/passgen", sessionStore.LoadSession(http.HandlerFunc(examplePassgenHandler), true))
	mux.Handle("/example/database", sessionStore.LoadSession(http.HandlerFunc(exampleDatabaseHandler), true))

	mux.Handle("/auth/login", sessionStore.LoadSession(http.HandlerFunc(loginHandler), false))
	mux.Handle("/auth/logout", sessionStore.LoadSession(http.HandlerFunc(logoutHandler), true))
	mux.Handle("/auth/logoutall", sessionStore.LoadSession(http.HandlerFunc(logoutAllHandler), true))

	mux.Handle("/account/sessions", sessionStore.LoadSession(http.HandlerFunc(accountSessionHandler), true))
	mux.Handle("/account/info", sessionStore.LoadSession(http.HandlerFunc(accountInfoHandler), true))

	mux.Handle("/api/identity", sessionStore.LoadSession(http.HandlerFunc(apiIdentityHandler), true))
	mux.Handle("/api/clientfetch", sessionStore.LoadSession(http.HandlerFunc(apiClientFetchHandler), true))

	log.Println("Mapped dynamic routes")
}
