package api

import (
	"net/http"
	"saral/auth"
	"saral/middleware"

	"log"
)

type LoginInfo struct {
	Username string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo LoginInfo

	err := ApiReadJSON(w, r, &loginInfo)
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

	ApiWritePlaintext(w, encrypted)
}
