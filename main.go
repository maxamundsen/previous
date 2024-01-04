package main

import (
	"errors"
	"fmt"
	"gohttp/constants"
	"net/http"
	"os"
)

var mux *http.ServeMux

func main() {
	mux = http.NewServeMux()

	fmt.Println("[Go HTTP Server Test]")

	InitMiddleware()
	MapStaticAssets(false)
	MapDynamicRoutes()

	fmt.Println("-> Listening on http://" + constants.HttpPort)

	err := http.ListenAndServe(constants.HttpPort, mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("-> [ERROR] Server closed\n")
	} else if err != nil {
		fmt.Printf("-> [ERROR] Starting server: %s\n", err)
		os.Exit(1)
	}
}
