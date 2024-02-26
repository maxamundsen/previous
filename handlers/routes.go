package handlers

import (
	"log"
	"net/http"
	"os"
)

// Dynamic routes are defined here
func MapDynamicRoutes(mux *http.ServeMux) {
	// root (index) handler
	mux.HandleFunc("/", indexHandler)

	// example handlers
	mux.HandleFunc("/example", sessionStore.LoadSession(exampleHandler, true))
	mux.HandleFunc("/example/upload", sessionStore.LoadSession(exampleUploadHandler, true))
	mux.HandleFunc("/example/counter", sessionStore.LoadSession(exampleCounterHandler, true))
	mux.HandleFunc("/example/passgen", sessionStore.LoadSession(examplePassgenHandler, true))
	mux.HandleFunc("/example/database", sessionStore.LoadSession(exampleDatabaseHandler, true))
	mux.HandleFunc("/example/adduser", sessionStore.LoadSession(exampleAdduserHandler, true))
	mux.HandleFunc("/example/mail", sessionStore.LoadSession(exampleMailHandler, true))
	mux.HandleFunc("/example/fetch", sessionStore.LoadSession(exampleFetchHandler, true))

	// auth handlers
	mux.HandleFunc("/auth/login", sessionStore.LoadSession(loginHandler, false))
	mux.HandleFunc("/auth/logout", sessionStore.LoadSession(logoutHandler, true))
	mux.HandleFunc("/auth/logoutall", sessionStore.LoadSession(logoutAllHandler, true))

	// account handlers
	mux.HandleFunc("/account/sessions", sessionStore.LoadSession(accountSessionHandler, true))
	mux.HandleFunc("/account/info", sessionStore.LoadSession(accountInfoHandler, true))

	// api handlers
	mux.HandleFunc("/api/identity", sessionStore.LoadSession(apiIdentityHandler, true))

	log.Println("Mapped dynamic routes")
}

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
