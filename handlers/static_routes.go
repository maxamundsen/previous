package handlers

import (
	"log"
	"os"
	"net/http"
)

// helper function to map endpoints for static assets.
func MapStaticAssets(mux *http.ServeMux) {
	dirPath := "wwwroot"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		log.Fatal(err)
	}

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("wwwroot/assets"))))
	mux.Handle("/favicon.ico", http.FileServer(http.Dir("wwwroot")))
	log.Println("Mapped static assets")
}
