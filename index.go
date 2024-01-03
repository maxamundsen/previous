package main

import (
	"io"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>This is the index page</h1>")
}
