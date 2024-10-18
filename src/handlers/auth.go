package auth

import (
	"webdawgengine/auth"
	"webdawgengine/config"
	"webdawgengine/middleware"
	"webdawgengine/views"
	"log"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	viewBase := views.ViewBase{
		Title: "Login",
	}

	ctx := views.NewViewContext(r, nil, &viewBase, nil)

	if r.Method == http.MethodGet {
		views.RenderWebpage(w, "auth/login", ctx)
	} else if r.Method == http.MethodPost {
		r.ParseForm()

		username := r.FormValue("username")
		password := r.FormValue("password")
		rememberMe, _ := strconv.ParseBool(r.FormValue("rememberMe"))

		user, authResult := auth.Authenticate(username, password)

		if !authResult {
			log.Println("Failed login attempt. Username: " + username)
			viewBase.ErrorMsg = "Username or password incorrect."
			views.RenderWebpage(w, "auth/login", ctx)
			return
		}

		log.Println("Successful login. Username: " + username)

		identity := auth.GenerateIdentity(&user, rememberMe)
		middleware.PutIdentityCookie(w, r, identity)

		params := r.URL.Query()
		location := params.Get("redirect")

		if len(params["redirect"]) > 0 {
			http.Redirect(w, r, location, http.StatusFound)
			return
		}

		if auth.IsClaimTruthy(identity.Claims, auth.U_EMPLOYEE) {
			http.Redirect(w, r, config.IDENTITY_EMPLOYEE_DEFAULT_PATH, http.StatusFound)
		} else {
			http.Redirect(w, r, config.IDENTITY_DEFAULT_PATH, http.StatusFound)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteIdentityCookie(w, r)

	http.Redirect(w, r, config.IDENTITY_LOGIN_PATH, http.StatusFound)
}
