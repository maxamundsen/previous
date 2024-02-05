package main

import (
	"fmt"
	"webdawgengine/build"
	"webdawgengine/config"
	"webdawgengine/database"
	"webdawgengine/handlers"
	"log"
	"net/http"
)

// Entry point for the application, initializes package globals
// such as the database connections, http multiplexer, config, etc.
func main() {
	fmt.Println("WebDawgEngine Initialized")

	if build.DEVEL {
		fmt.Println("*DEVELOPMENT BUILD")
	} else {
		fmt.Println("*RELEASE BUILD")
	}

	config.ReadConfiguration()
	config := config.GetConfiguration()

	database.InitializeDb(config.ConnectionString)
	handlers.SessionInit(config.CookieExpiryDays)
	mux := http.NewServeMux()

	handlers.MapStaticAssets(mux)
	handlers.MapDynamicRoutes(mux)

	log.Println("Listening on http://" + config.Host)

	err := http.ListenAndServe(config.Host, mux)

	if err != nil {
		log.Fatal(err)
	}
}
