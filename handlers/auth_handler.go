package handlers

import (
	"gohttp/auth"
	"gohttp/views"
	"net/http"
	"strconv"
	"crypto/rand"
	"time"
	"math/big"
)

type loginPage struct {
	Base views.ViewBase
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// time attack partial mitigation
	// this does not fully prevent them, but should make it very unlikely when
	// paired with a 'max login attempt per ip/minute/fingerprint'
	// https://security.stackexchange.com/questions/96489/can-i-prevent-timing-attacks-with-random-delays/96493#96493
	// https://www.reddit.com/r/PHP/comments/kn6ezp/have_you_secured_your_signup_process_against_a/
	randomSeconds, _ := rand.Int(rand.Reader, big.NewInt(1000))
	randomDuration := time.Duration(randomSeconds.Int64()) * time.Millisecond
	
	viewData := make(map[string]string)
	
	viewData["Title"] = "Login"
	
	base := views.NewViewBase(nil, viewData);
	
	pageData := loginPage {
		base,
	}
	
	if r.Method == http.MethodGet {
		views.RenderTemplate(w, "login", pageData)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		rememberMe, _ := strconv.ParseBool(r.FormValue("rememberMe"))
	
		// add random time between 0-1 seconds to response
		time.Sleep(randomDuration)
		
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
			
			pageData.Base.ViewData["Error"] = "Username or password incorrect."
			views.RenderTemplate(w, "login", pageData)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionStore.DeleteSession(w, r)
	
	storeBase := sessionStore.GetBase()
	http.Redirect(w, r, storeBase.LoginPath, http.StatusFound)
}
