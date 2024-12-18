package main

import (
	"fmt"
	"log"
	"net/http"
	"saral/build"
	"saral/config"
	"saral/database"

	"saral/pages/docs"
)

// Entry point for the application, initializes package globals
// such as the database connections, http multiplexer, config, etc.
func main() {
	fmt.Println("Saral V2")

	if build.DEBUG {
		fmt.Println("DEBUG BUILD")
	} else {
		fmt.Println("RELEASE BUILD")
	}

	config.LoadConfig()

	docs.RegisterDocumentation()
	database.Init()

	// create http multiplexer, map routes
	mux := http.NewServeMux()

	mapPageRoutes(mux)
	mapApiRoutes(mux)

	log.Println("Listening on http://" + config.GetConfig().Host + ":" + config.GetConfig().Port)

	serveErr := http.ListenAndServe(config.GetConfig().Host+":"+config.GetConfig().Port, mux)

	if serveErr != nil {
		log.Fatal(serveErr)
	}
}
