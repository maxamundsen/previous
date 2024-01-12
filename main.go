package main

import (
	"fmt"
	"gohttp/constants"
	"gohttp/constants/build"
	"gohttp/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go HTTP Server Test")

	if build.DEVEL {
		fmt.Println("DEVELOPMENT MODE ENABLED")
		fmt.Println("Note: View templates are NOT embedded in devel mode")
	} else {
		fmt.Println("PRODUCTION BUILD")
	}

	// Create in-memory session store
	handlers.SessionInit()

	// Create http multiplexer
	mux := http.NewServeMux()

	if build.EMBED {
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
