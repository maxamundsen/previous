package handlers

import (
	"gohttp/auth"
	"gohttp/views"
	"net/http"
)

func passGenHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	if identity.Claims["CanGeneratePasswords"] != "true" {
		return
	}

	val1 := r.FormValue("val1")

	var password string

	if val1 == "" {
		password = "empty"
	} else {
		password, _ = auth.HashPassword(val1)
	}

	viewData := make(map[string]interface{})
	viewData["Password"] = password

	base := views.NewViewModel(identity, viewData)

	views.RenderTemplate(w, "passgen", base)
}
