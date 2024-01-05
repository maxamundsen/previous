package main

import (
	"fmt"
	"gohttp/constants"
	"log"
	"net/http"
)

var mux *http.ServeMux

func main() {
	mux = http.NewServeMux()

	fmt.Println("[Go HTTP Server Test]")
	fmt.Println("")

	InitMiddleware()
	MapStaticAssets(false)
	MapDynamicRoutes()

	log.Println("Listening on http://" + constants.HttpPort)

	err := http.ListenAndServe(constants.HttpPort, mux)

	if err != nil {
		log.Fatal(err)
	}
}
