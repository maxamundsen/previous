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
			sess := &auth.AuthSession{}
			sess.IsAuthenticated = true
			sess.RememberMe = rememberMe
			sess.Role = "Administrator"
			sess.Username = username
	
			sessionStore.PutSession(w, r, sess)
	
			params := r.URL.Query()
			location := params.Get("redirect")
	
			if len(params["redirect"]) > 0 {
				http.Redirect(w, r, location, http.StatusFound)
				return
			}
	
			base := sessionStore.GetBase()
			http.Redirect(w, r, base.LoginPath, http.StatusFound)
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
