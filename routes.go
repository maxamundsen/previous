package main

import (
	"embed"
	"net/http"
	"fmt"
)

//go:embed assets
var staticAssets embed.FS

//go:embed views
var viewTemplates embed.FS

// wrapper function to simply session init process
// when requireAuth is true, routes will 401, or redirect to login if
// a valid session is not found
func handleWithSession(route string, function http.HandlerFunc, requireAuth bool) {
	// attach auth session middleware to provided route
	http.Handle(route, globalSession.LoadSession(http.HandlerFunc(function), requireAuth))
}

func mapStaticAssets() {	
	staticSrv := http.FS(staticAssets)
	fs := http.FileServer(staticSrv)
	http.Handle("/assets/", fs)
	
	fmt.Println("-> Mapped static assets [EMBED=TRUE]")
}

func mapDynamicRoutes() {
	handleWithSession("/", IndexHandler, true)
	handleWithSession("/hello", HelloHandler, true)
	handleWithSession("/login", LoginHandler, false)
	
	fmt.Println("-> Mapped dynamic routes")
}
