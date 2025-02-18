package main

import (
	"fmt"
	"log"
	"net/http"
	"previous/config"
	"previous/pages"
	"previous/repository"
	"previous/security"
	"previous/tasks"
)

func main() {
	fmt.Println("Previous: A powerful web codebase.\n")

	if config.DEBUG {
		fmt.Println("DEBUG BUILD")
	} else {
		fmt.Println("RELEASE BUILD")
	}

	config.Init()
	security.Init()
	repository.Init()
	pages.Init()
	tasks.Init()

	mux := http.NewServeMux()
	mapRoutes(mux)

	log.Println("Mapped HTTP routes")
	log.Println("Listening on http://" + config.GetConfig().Host + ":" + config.GetConfig().Port)

	serveErr := http.ListenAndServe(config.GetConfig().Host+":"+config.GetConfig().Port, mux)

	if serveErr != nil {
		log.Fatal(serveErr)
	}
}
