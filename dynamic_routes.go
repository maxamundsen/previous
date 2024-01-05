package main

import (
	"gohttp/handlers"
	"log"
	"net/http"
)

// Dynamic routes are defined here
func MapDynamicRoutes() {
	mux.HandleFunc("/", handlers.IndexHandler)
	handleFuncWithSession("/test", handlers.TestHandler, true)
	handleFuncWithSession("/login", handlers.LoginHandler, false)
	handleFuncWithSession("/logout", handlers.LogoutHandler, true)

	log.Println("Mapped dynamic routes")
}

// wrapper function for working with session authentication middleware
func handleFuncWithSession(route string, handler http.HandlerFunc, requireAuth bool) {
	mux.Handle(route, handlers.MemorySession.LoadSession(http.HandlerFunc(handler), requireAuth))
}
