package api

import (
	"net/http"
	"webdawgengine/middleware"
)

func AccountController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	ApiWriteJSON(w, identity.User)
}
