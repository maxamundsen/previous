package main

import (
	"fmt"
	"gohttp/build"
	"gohttp/config"
	"gohttp/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go HTTP Server Test")

	if build.DEVEL {
		fmt.Println("*DEVELOPMENT BUILD")
	} else {
		fmt.Println("*RELEASE BUILD")
	}

	// Read "config.json" file
	config.ReadConfiguration()
	config := config.GetConfiguration()

	// Create in-memory session store
	handlers.SessionInit()

	// Create http multiplexer
	mux := http.NewServeMux()

	// when the `embed` build tag is set, static assets will be
	// embedded in the binary, and served from the embedded filesystem
	if build.EMBED {
		handlers.MapStaticAssetsEmbed(mux, &staticAssets)
	} else {
		handlers.MapStaticAssets(mux)
	}

	handlers.MapDynamicRoutes(mux)

	log.Println("Listening on http://" + config.Host)

	err := http.ListenAndServe(config.Host, mux)

	if err != nil {
		log.Fatal(err)
	}
}
