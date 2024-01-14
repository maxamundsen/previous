package handlers

import (
	"gohttp/views"
	"net/http"
)

func accountSessionHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	identityList := sessionStore.GetAllIdentities(identity)

	viewData := make(map[string]interface{})
	viewData["Title"] = "Login Sessions"
	viewData["IdentityList"] = identityList

	model := views.NewViewModel(identity, viewData)
	views.RenderTemplate(w, "account_sessions", model)
}
