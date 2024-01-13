package handlers

import (
	"gohttp/auth"
	"gohttp/views"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	val1 := r.FormValue("val1")

	var password string

	if val1 == "" {
		password = "empty"
	} else {
		password, _ = auth.HashPassword(val1)
	}

	viewData := make(map[string]interface{})

	viewData["Title"] = "Test Page"
	viewData["Password"] = password

	base := views.NewViewModel(identity, viewData)

	views.RenderTemplate(w, "test", base)
}
