package api

import (
	"net/http"
	"previous/middleware"
)

func AccountApiPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	ApiWriteJSON(w, identity.User)
}
