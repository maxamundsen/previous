package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"webdawgengine/auth"
	"webdawgengine/views"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})

	viewData["title"] = "Login"

	model := views.NewViewModel(nil, viewData)

	if r.Method == http.MethodGet {
		views.RenderWebpage(w, "login", model)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		rememberMe, _ := strconv.ParseBool(r.FormValue("rememberMe"))

		id, err := auth.Authenticate(username, password)

		if err != nil {
			log.Println("Failed login attempt. Username: " + username)
			model.ViewData["error"] = "Username or password incorrect."
			views.RenderWebpage(w, "login", model)
			return
		}

		log.Println("Successful login. Username: " + username)

		id.RememberMe = rememberMe
		id.IpAddr = r.RemoteAddr
		id.UserAgent = r.UserAgent()
		id.LoginTime = time.Now()
		sessionStore.PutSession(w, r, id)

		params := r.URL.Query()
		location := params.Get("redirect")

		if len(params["redirect"]) > 0 {
			http.Redirect(w, r, location, http.StatusFound)
			return
		}

		storeBase := sessionStore.GetBase()
		http.Redirect(w, r, storeBase.DefaultPath, http.StatusFound)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	id := sessionStore.GetIdentityFromCtx(r)

	sessionStore.DeleteSession(w, r)

	storeBase := sessionStore.GetBase()
	log.Println("Successful logout. Email: " + id.Email)

	http.Redirect(w, r, storeBase.LoginPath, http.StatusFound)
}

func logoutAllHandler(w http.ResponseWriter, r *http.Request) {
	id := sessionStore.GetIdentityFromCtx(r)
	sessionStore.DeleteAllByEmail(w, r, id)

	log.Println("Successful all-session logout for Email: " + id.Email)

	storeBase := sessionStore.GetBase()
	http.Redirect(w, r, storeBase.LoginPath, http.StatusFound)
}
