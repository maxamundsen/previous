package handlers

import (
	"gohttp/auth"
	"gohttp/views"
	"gohttp/data"
	"net/http"
	"strconv"
)

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)

	viewData := make(map[string]interface{})
	viewData["Title"] = "Example Page"

	model := views.NewViewModel(identity, viewData)
	views.RenderTemplate(w, "example", model)
}

func exampleCounterHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})

	formVal := r.FormValue("number")

	num, _ := strconv.Atoi(formVal)

	viewData["Number"] = num + 1

	model := views.NewViewModel(nil, viewData)
	views.RenderTemplate(w, "example_counter", model)
}

func examplePassgenHandler(w http.ResponseWriter, r *http.Request) {
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

	model := views.NewViewModel(identity, viewData)

	views.RenderTemplate(w, "example_passgen", model)
}

func exampleDatabaseHandler(w http.ResponseWriter, r *http.Request) {
	identity := sessionStore.GetIdentityFromCtx(r)
	
	users := data.FetchUsers()
	
	viewData := make(map[string]interface{})
	viewData["Users"] = users
	
	model := views.NewViewModel(identity, viewData)

//TODO UPDATE VIEW 
	views.RenderTemplate(w, "example_passgen", model)
}