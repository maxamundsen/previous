package main

import (
	"errors"
	"fmt"
	"gohttp/auth"
	"gohttp/constants"
	"net/http"
	"os"
)

var memorySession auth.MemorySessionStore

func main() {
	fmt.Println("[Go HTTP Server Test]")

	memorySession.InitStore("AuthenticationCookie", constants.CookieExpiryTime, true, "/login", "/logout", "/test")

	mapStaticAssets()
	mapDynamicRoutes()

	fmt.Println("-> Listening on http://" + constants.HttpPort)

	err := http.ListenAndServe(constants.HttpPort, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("-> [ERROR] Server closed\n")
	} else if err != nil {
		fmt.Printf("-> [ERROR] Starting server: %s\n", err)
		os.Exit(1)
	}
}
