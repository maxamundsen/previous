package api

import (
	"net/http"
	"previous/middleware"
)

// @Identity
// @Protected
func AccountPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	ApiWriteJSON(w, identity.User)
}
