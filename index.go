package main

import (
	"fmt"
	"io"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got index request\n")
	io.WriteString(w, "<h1>This is the index page</h1>")
}