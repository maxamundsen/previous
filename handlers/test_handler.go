package handlers

import (
	"gohttp/views"
	"net/http"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	viewData := make(map[string]interface{})
	viewData["Title"] = "Example Page"

	base := views.NewViewModel(identity, viewData)
	views.RenderTemplate(w, "example", base)
}
