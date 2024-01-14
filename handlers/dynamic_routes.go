package handlers

import (
	"log"
	"net/http"
)

// Dynamic routes are defined here
func MapDynamicRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)

	mux.Handle("/test", sessionStore.LoadSession(http.HandlerFunc(testHandler), true))

	mux.Handle("/auth/login", sessionStore.LoadSession(http.HandlerFunc(loginHandler), false))
	mux.Handle("/auth/logout", sessionStore.LoadSession(http.HandlerFunc(logoutHandler), true))
	mux.Handle("/auth/logoutall", sessionStore.LoadSession(http.HandlerFunc(logoutAllHandler), true))

	mux.Handle("/account/sessions", sessionStore.LoadSession(http.HandlerFunc(accountSessionHandler), true))

	mux.Handle("/api/test", sessionStore.LoadSession(http.HandlerFunc(apiTestHandler), true))
	mux.Handle("/api/user", sessionStore.LoadSession(http.HandlerFunc(apiUserHandler), true))
	mux.Handle("/api/clientfetch", sessionStore.LoadSession(http.HandlerFunc(apiClientFetchHandler), true))

	log.Println("Mapped dynamic routes")
}
