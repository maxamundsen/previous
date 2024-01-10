package handlers

import (
	"gohttp/auth"
	"gohttp/views"
	"net/http"
)

type pageData struct {
	User   	 *auth.Identity
	Title    string
	Password string
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	user := sessionStore.GetIdentityFromCtx(r)

	val1 := r.FormValue("val1")

	var password string

	if val1 == "" {
		password = "empty"
	} else {
		password, _ = auth.HashPassword(val1)
	}

	pageData := pageData{
		user,
		"Title for page",
		password,
	}

	views.RenderTemplate(w, "test", pageData)
}
