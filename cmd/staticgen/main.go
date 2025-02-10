package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"previous/pages/app"
	. "previous/middleware"
)


func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/app/dashboard", LoadIdentity(LoadSessionFromCookie(app.DashboardPage), true))

	req, _ := http.NewRequest("GET", "/app/dashboard", nil)

	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)

	fmt.Printf("%s", rr.Body.String())
}