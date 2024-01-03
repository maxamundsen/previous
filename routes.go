package main

import (
	"net/http"
	"embed"
)

//go:embed assets
var staticAssets embed.FS

//go:embed views
var viewTemplates embed.FS

func handleWithSession(route string, function http.HandlerFunc) {
	http.Handle(route, globalSession.LoadSession(http.HandlerFunc(function)))
}

func mapStaticAssets() {
	staticSrv := http.FS(staticAssets)
	fs := http.FileServer(staticSrv)
	http.Handle("/assets/", fs)
}

func mapDynamicRoutes() {
	http.HandleFunc("/", IndexHandler)
	handleWithSession("/hello", HelloHandler)
	handleWithSession("/login", LoginHandler)
}