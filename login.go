package main

import (
	"net/http"
	"html/template"
	"gohttp/session"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFS(viewTemplates, "views/base.html", "views/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		
		if username == "admin" && password == "admin" {
			sess := &session.AuthSession{}
			sess.Role = "Administrator"
			sess.Username = username
			ss.PutSession(w, r, sess)
			http.Redirect(w, r, "/hello", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			tmpl := template.Must(template.ParseFS(viewTemplates, "views/error.html"))
			tmpl.Execute(w, "ERROR FUCK YOU") 
		}
	}
}