package main

import (
	"fmt"
	"log"
	"net/http"
	"previous/.metagen/router"
	"previous/config"
	"previous/preload"
)

// WARNING:
// DO NOT INITIALIZE GLOBAL STATE HERE!!
// Global state should be initialized by the `preload` package.
// This is to ensure that `metagen`, the codebase metaprogram, can arbitrarily execute
// code at compile time.
func main() {
	fmt.Println("Previous: A powerful web codebase.\n")

	if config.DEBUG {
		fmt.Println("DEBUG BUILD")
	} else {
		fmt.Println("RELEASE BUILD")
	}

	// Preload module handles global state because Golang import system is weird...
	preload.PreloadInit(preload.PreloadOptionsAll())

	router.MetagenAutoRouter(preload.HttpMux)
	mapManualRoutes(preload.HttpMux)

	log.Println("Mapped HTTP routes")

	log.Println("Listening on http://" + config.GetConfig().Host + ":" + config.GetConfig().Port)

	serveErr := http.ListenAndServe(config.GetConfig().Host+":"+config.GetConfig().Port, preload.HttpMux)

	if serveErr != nil {
		log.Fatal(serveErr)
	}
}
