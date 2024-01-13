package handlers

import (
	"crypto/rand"
	"gohttp/auth"
	"gohttp/views"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	viewData := make(map[string]interface{})

	viewData["Title"] = "Login"

	base := views.NewViewModel(nil, viewData)

	if r.Method == http.MethodGet {
		views.RenderTemplate(w, "login", base)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		rememberMe, _ := strconv.ParseBool(r.FormValue("rememberMe"))

		// time attack partial mitigation
		// adds up to 0.5 seconds to the response time

		// this technically does not prevent a time attack, since there is still time variance without the randomness added.
		// you could theoretically take an average of a 'valid user; incorrect password' vs 'invalid user' response times
		// to figure out if a user exists, but you would need a lot of data to do that.
		// this should make it *extremely* unlikely to do when paired with 'n login attempt per ip/minute/fingerprint'
		// since you would need way more than `n` login attempts to collect an average

		// https://security.stackexchange.com/questions/96489/can-i-prevent-timing-attacks-with-random-delays/96493#96493
		// https://www.reddit.com/r/PHP/comments/kn6ezp/have_you_secured_your_signup_process_against_a/

		randomSeconds, _ := rand.Int(rand.Reader, big.NewInt(500))
		randomDuration := time.Duration(randomSeconds.Int64()) * time.Millisecond

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

			base.ViewData["Error"] = "Username or password incorrect."
			views.RenderTemplate(w, "login", base)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sessionStore.DeleteSession(w, r)

	storeBase := sessionStore.GetBase()
	http.Redirect(w, r, storeBase.LoginPath, http.StatusFound)
}
