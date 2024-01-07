package handlers

import (
	"log"
	"net/http"
	"gohttp/auth"
)

// Dynamic routes are defined here
func MapDynamicRoutesWithMemoryStore(mux *http.ServeMux, store *auth.MemorySessionStore) {
	mux.HandleFunc("/", IndexHandler)
	// handleFuncWithSession("/test", mux, store, TestHandler, true)
	handleFuncWithSession("/login", mux, store, LoginHandler, false)
	// handleFuncWithSession("/logout", mux, store, LogoutHandler, true)

	log.Println("Mapped dynamic routes")
}

// wrapper function for working with session authentication middleware
func handleFuncWithSession(route string, 
                           mux *http.ServeMux,
                           store *auth.MemorySessionStore, 
                           handler http.Handler, 
                           requireAuth bool) {
	mux.Handle(route, store.LoadSession(handler, requireAuth))
}
