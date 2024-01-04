package main

import (
	"embed"
	"fmt"
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
	var staticSrv http.FileSystem

	if embed {
		staticSrv = http.FS(staticAssets)
	} else {
		staticSrv = http.Dir("../assets")
	}

	fs := http.FileServer(staticSrv)
	http.Handle("/assets/", fs)

	fmt.Println("-> Mapped static assets")
}

func MapDynamicRoutes() {
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.Handle("/test", handlers.MemorySession.LoadSession(http.HandlerFunc(handlers.TestHandler), true))
	mux.Handle("/login", handlers.MemorySession.LoadSession(http.HandlerFunc(handlers.LoginHandler), false))
	mux.Handle("/logout", handlers.MemorySession.LoadSession(http.HandlerFunc(handlers.LogoutHandler), true))

	fmt.Println("-> Mapped dynamic routes")
}
