package handlers

import (
	"gohttp/auth"
	"gohttp/views"
	"net/http"
	"strconv"
)


func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		views.RenderTemplate(w, "login", nil)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		rememberMe, _ := strconv.ParseBool(r.FormValue("rememberMe"))
	
		if username == "admin" && password == "admin" {
			id := &auth.Identity{}
			id.IsAuthenticated = true
			id.RememberMe = rememberMe
			id.Role = "Administrator"
			id.Username = username
	
			sessionStore.PutSession(w, r, id)
	
			params := r.URL.Query()
			location := params.Get("redirect")
	
			if len(params["redirect"]) > 0 {
				http.Redirect(w, r, location, http.StatusFound)
				return
			}
	
			storeBase := sessionStore.GetBase()
			http.Redirect(w, r, storeBase.DefaultPath, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			views.RenderTemplate(w, "error", nil)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionStore.DeleteSession(w, r)
	
	storeBase := sessionStore.GetBase()
	http.Redirect(w, r, storeBase.LoginPath, http.StatusFound)
}
