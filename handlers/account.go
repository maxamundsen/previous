package handlers

import (
	"net/http"
	"webdawgengine/views"
)

func accountSessionHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	identityList := sessionStore.GetAllIdentities(identity)

	viewData := make(map[string]interface{})
	viewData["title"] = "Login Sessions"
	viewData["identity_list"] = identityList

	model := views.NewViewModel(identity, viewData)
	views.RenderWebpage(w, "account_sessions", model)
}

func accountInfoHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	identityList := sessionStore.GetAllIdentities(identity)

	viewData := make(map[string]interface{})
	viewData["title"] = "Account Info"
	viewData["identity_list"] = identityList

	model := views.NewViewModel(identity, viewData)
	views.RenderWebpage(w, "account_info", model)
}
