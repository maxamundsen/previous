package main

import (
	"log"
	"embed"
	"gohttp/constants"
	"gohttp/handlers"
	"net/http"
)

//go:embed assets
var staticAssets embed.FS

func InitMiddleware() {
	handlers.MemorySession.InitStore("AuthenticationCookie", constants.CookieExpiryTime, true, "/login", "/logout", "/test")
}

func MapStaticAssets(embed bool) {
	if embed {
		mux.Handle("/assets/", http.FileServer(http.FS(staticAssets)))
	} else {
		mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	}

	log.Printf("Mapped static assets [embed: %t] \n", embed)
}

func MapDynamicRoutes() {
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.Handle("/test", handlers.MemorySession.LoadSession(http.HandlerFunc(handlers.TestHandler), true))
	mux.Handle("/login", handlers.MemorySession.LoadSession(http.HandlerFunc(handlers.LoginHandler), false))
	mux.Handle("/logout", handlers.MemorySession.LoadSession(http.HandlerFunc(handlers.LogoutHandler), true))

	log.Println("Mapped dynamic routes")
}
