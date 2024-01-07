package main

import (
	"embed"
	"fmt"
	"gohttp/constants"
	"gohttp/handlers"
	"gohttp/auth"
	"log"
	"net/http"
)

//go:embed wwwroot/favicon.ico
var content embed.FS

//go:embed wwwroot/assets
var staticAssets embed.FS

func main() {
	// Create in-memory session store
	store := &auth.MemorySessionStore{}
	store.InitStore("AuthenticationCookie", constants.CookieExpiryTime, true, "/login", "/logout", "/test")
	
	// Create http multiplexer
	mux := http.NewServeMux()

	fmt.Printf("[Go HTTP Server Test]\n\n")

	handlers.SessionInit()

	if constants.UseEmbed {
		handlers.MapStaticAssetsEmbed(mux, &staticAssets)
	} else {
		handlers.MapStaticAssets(mux)
	}

	handlers.MapDynamicRoutesWithMemoryStore(mux, store)

	log.Println("Listening on http://" + constants.HttpPort)

	err := http.ListenAndServe(constants.HttpPort, mux)

	if err != nil {
		log.Fatal(err)
	}
}
