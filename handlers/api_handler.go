package handlers

import (
	"encoding/json"
	"net/http"
)

// Json endpoints for an API can be easily written with a handler.
// In order to parse json, you must specify some information about the output json
// via decorators.
// structs can be automatically generated from json using: https://mholt.github.io/json-to-go/

// This will print information about the current Identity
// via session access
func apiIdentityHandler(w http.ResponseWriter, r *http.Request) {
	user := sessionStore.GetIdentityFromCtx(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}