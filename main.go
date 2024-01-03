package main

import (
	"errors"
	"fmt"
	"gohttp/session"
	"net/http"
	"os"
	"time"
)

var globalSession session.SessionStore[session.AuthSession]

func main() {
	const expiry time.Duration = time.Duration(time.Hour*24*7) // 7 days
	const location string = "localhost:8080"

	fmt.Println("[Go HTTP Server Test]")

	globalSession.InitStore("AuthenticationCookie", expiry, true, "/login", "/hello")
	fmt.Println("-> Initializing in-memory session")

	mapStaticAssets()
	fmt.Println("-> Embeded assets and templates are ENABLED")
	
	mapDynamicRoutes()

	fmt.Println("-> Listening on http://" + location)
	err := http.ListenAndServe(location, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
