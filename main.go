package main

import (
	"fmt"
	"gohttp/constants"
	"gohttp/handlers"
	"log"
	"net/http"
)

var mux *http.ServeMux

var useEmbed bool = true

func main() {
	mux = http.NewServeMux()

	fmt.Println("[Go HTTP Server Test]")
	fmt.Println("")

	handlers.MemorySession.InitStore("AuthenticationCookie", constants.CookieExpiryTime, true, "/login", "/logout", "/test")

	MapStaticAssets()
	MapDynamicRoutes()

	log.Println("Listening on http://" + constants.HttpPort)

	err := http.ListenAndServe(constants.HttpPort, mux)

	if err != nil {
		log.Fatal(err)
	}
}
