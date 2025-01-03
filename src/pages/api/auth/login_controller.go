package auth

import (
	"net/http"
	"previous/auth"
	"previous/middleware"
	"previous/pages/api"

	"log"
)

type LoginInfo struct {
	Username string
	Password string
}

// @Identity
func LoginController(w http.ResponseWriter, r *http.Request) {
	var loginInfo LoginInfo

	err := api.ApiReadJSON(w, r, &loginInfo)
	if err != nil {
		println(err)
	}

	userid, authResult := auth.Authenticate(loginInfo.Username, loginInfo.Password)
	if !authResult {
		log.Println("Failed login attempt via API call. Username: " + loginInfo.Username)
		http.Error(w, "Invalid login information", http.StatusUnauthorized)
		return
	}

	identity := auth.NewIdentity(userid, false)

	encrypted, err := middleware.EncryptIdentity(identity)
	if err != nil {
		http.Error(w, "An error occured while generating the api token.", http.StatusUnauthorized)
		return
	}

	api.ApiWritePlaintext(w, encrypted)
}
