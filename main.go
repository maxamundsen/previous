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

	config.ParseConfigFile()
	database.InitializeDb()
	handlers.SessionInit()

	mux := http.NewServeMux()

	handlers.MapStaticAssets(mux)
	handlers.MapDynamicRoutes(mux)

	log.Println("Listening on http://" + config.GetHost())

	err := http.ListenAndServe(config.GetHost(), mux)

	if err != nil {
		log.Fatal(err)
	}
}
