package main

import (
	"fmt"
	"gohttp/constants"
	"gohttp/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("[Go HTTP Server Test]\n\n")

	// Create in-memory session store
	handlers.SessionInit()

	// Create http multiplexer
	mux := http.NewServeMux()
	
	if constants.EMBED {
		handlers.MapStaticAssetsEmbed(mux, &staticAssets)
	} else {
		handlers.MapStaticAssets(mux)
	}

	handlers.MapDynamicRoutes(mux)

	log.Println("Listening on http://" + constants.HttpPort)

	err := http.ListenAndServe(constants.HttpPort, mux)

	if err != nil {
		log.Fatal(err)
	}
}
