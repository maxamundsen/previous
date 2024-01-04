package handlers

import (
	"gohttp/auth"
	"html/template"
	"net/http"
	"strconv"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("views/base.html", "views/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		rememberMe, _ := strconv.ParseBool(r.FormValue("rememberMe"))

		if username == "admin" && password == "admin" {
			sess := &auth.AuthSession{}
			sess.IsAuthenticated = true
			sess.RememberMe = rememberMe
			sess.Role = "Administrator"
			sess.Username = username

			MemorySession.PutSession(w, r, sess)

			params := r.URL.Query()
			location := params.Get("redirect")

			if len(params["redirect"]) > 0 {
				http.Redirect(w, r, location, http.StatusFound)
				return
			}

			http.Redirect(w, r, MemorySession.Base.LoginPath, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			tmpl := template.Must(template.ParseFiles("views/error.html"))
			tmpl.Execute(w, "LOGIN ERROR")
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	MemorySession.DeleteSession(r)
	http.Redirect(w, r, MemorySession.Base.LoginPath, http.StatusFound)
}
