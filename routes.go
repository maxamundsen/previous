package main

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed assets
var staticAssets embed.FS

//go:embed views
var viewTemplates embed.FS

func mapStaticAssets() {
	staticSrv := http.FS(staticAssets)
	fs := http.FileServer(staticSrv)
	http.Handle("/assets/", fs)

	fmt.Println("-> Mapped static assets [EMBED=TRUE]")
}

func mapDynamicRoutes() {
	http.HandleFunc("/", IndexHandler)
	http.Handle("/test", memorySession.LoadSession(http.HandlerFunc(TestHandler), true))
	http.Handle("/login", memorySession.LoadSession(http.HandlerFunc(LoginHandler), false))
	http.Handle("/logout", memorySession.LoadSession(http.HandlerFunc(LogoutHandler), true))

	fmt.Println("-> Mapped dynamic routes")
}
