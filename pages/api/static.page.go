package api

import (
	"net/http"
)

// @StaticAPI
func StaticApiPage(w http.ResponseWriter, r *http.Request) {
	type model struct {
		Field1 int
		Field2 string
		Field3 bool
	}

	data := model{
		Field1: 1,
		Field2: "Hello, world!",
		Field3: true,
	}

	ApiWriteJSON(w, data)
}
