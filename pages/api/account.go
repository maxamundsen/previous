package api

import (
	"net/http"
	"previous/middleware"
)

// @Identity
// @Protected
// @EnableCors
func AccountApiPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	ApiWriteJSON(w, identity.User)
}
