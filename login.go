package main

import (
	"gohttp/session"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	sess := globalSession.GetSessionFromCtx(r)
	globalSession.AuthorizeRoute(w, r, sess)
	
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFS(viewTemplates, "views/base.html", "views/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "admin" && password == "admin" {
			sess := &session.AuthSession{}
			sess.IsAuthenticated = true
			sess.Role = "Administrator"
			sess.Username = username
			globalSession.PutSession(w, r, sess)
			http.Redirect(w, r, "/hello", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			tmpl := template.Must(template.ParseFS(viewTemplates, "views/error.html"))
			tmpl.Execute(w, "LOGIN ERROR")
		}
	}
}
