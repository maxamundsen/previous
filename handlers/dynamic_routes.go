package handlers

import (
	"log"
	"net/http"
)

// Dynamic routes are defined here
func MapDynamicRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)
	mux.Handle("/login", sessionStore.LoadSession(http.HandlerFunc(loginHandler), false))
	mux.Handle("/logout", sessionStore.LoadSession(http.HandlerFunc(logoutHandler), true))
	mux.Handle("/test", sessionStore.LoadSession(http.HandlerFunc(testHandler), true))
	
	log.Println("Mapped dynamic routes")
}