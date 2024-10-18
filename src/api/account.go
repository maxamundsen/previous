package api

import (
	"net/http"
	"webdawgengine/database"
	"webdawgengine/middleware"
)

func AccountController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	user, _ := database.FetchUserById(identity.UserId)

	ApiWriteJSON(w, user)
}
