package handlers

import (
	"gohttp/auth"
	"gohttp/views"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type PageData struct {
	Person   Person
	Title    string
	Password string
	IsAuth   bool
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	sess := MemorySession.Base.GetSessionFromCtx(r)

	val1 := r.FormValue("val1")

	var password string

	if val1 == "" {
		password = "empty"
	} else {
		password, _ = auth.HashPassword(val1)
	}

	// if sess.Role != "Administrator" {

	// }

	harambe := Person{
		"Firstname",
		"Lastname",
		15,
	}

	pageData := PageData{
		harambe,
		"Title for page",
		password,
		sess.IsAuthenticated,
	}

	views.RenderTemplate(w, "test", pageData)
}
