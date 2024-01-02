package main

import (
	"errors"
	"embed"
	"fmt"
	"net/http"
	"os"
	"time"
	"gohttp/session"
)

var ss session.SessionStore[session.AuthSession]

//go:embed assets
var staticAssets embed.FS

//go:embed views
var viewTemplates embed.FS

func mapStaticAssets() {
	staticSrv := http.FS(staticAssets)
	fs := http.FileServer(staticSrv)
	http.Handle("/assets/", fs)
}

func mapDynamicRoutes() {
	http.HandleFunc("/", IndexHandler)
	http.Handle("/hello", ss.LoadSession(http.HandlerFunc(HelloHandler)))
	http.HandleFunc("/login", LoginHandler)
}

func main() {
	const port string = "8080"
	
	fmt.Println("[Go HTTP Server Test]")
	fmt.Println("-> Listening on :" + port)
	
	ss.InitStore("SessionID", time.Duration(time.Hour*24*7)) // 1 week
	fmt.Println("-> Initializing in-memory session")
	
	mapStaticAssets()
	mapDynamicRoutes()

	err := http.ListenAndServe("localhost:" + port, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
