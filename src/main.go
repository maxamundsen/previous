package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"webdawgengine/build"
	"webdawgengine/config"
	"webdawgengine/database"
)

// Entry point for the application, initializes package globals
// such as the database connections, http multiplexer, config, etc.
func main() {
	fmt.Println("WebDawgEngine V2")

	if build.DEBUG {
		fmt.Println("DEBUG BUILD")
	} else {
		fmt.Println("RELEASE BUILD")
	}

	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config.LoadConfig(configFile)

	database.Init()

	// create http multiplexer, map routes
	mux := http.NewServeMux()

	mapPageRoutes(mux)
	mapApiRoutes(mux)

	log.Println("Listening on http://" + config.Config.Host + ":" + config.Config.Port)

	serveErr := http.ListenAndServe(config.Config.Host+":"+config.Config.Port, mux)

	if serveErr != nil {
		log.Fatal(serveErr)
	}
}
