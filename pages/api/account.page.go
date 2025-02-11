package api

import (
	"net/http"
	"previous/middleware"
)

// @Identity
// @Protected
// @EnableCors
func AccountPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	ApiWriteJSON(w, identity.User)
}
