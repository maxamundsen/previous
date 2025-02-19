package api

import (
	"net/http"
	"previous/middleware"
)

func AccountApiHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	ApiWriteJSON(w, identity.User)
}
