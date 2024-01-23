package main

import (
	"fmt"
	"gohttp/build"
	"gohttp/config"
	"gohttp/data"
	"gohttp/handlers"
	"log"
	"net/http"
)

// Entry point for the application, initializes package globals
// such as the database connections, http multiplexer, config, etc.
func main() {
	fmt.Println("Go HTTP Server Test")

	if build.DEVEL {
		fmt.Println("*DEVELOPMENT BUILD")
	} else {
		fmt.Println("*RELEASE BUILD")
	}

	config.ReadConfiguration()
	config := config.GetConfiguration()

	data.InitializeDb(config.ConnectionString)
	handlers.SessionInit(config.CookieExpiryDays)
	mux := http.NewServeMux()

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
