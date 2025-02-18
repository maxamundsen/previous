package auth

import (
	"net/http"
	"previous/config"
	"previous/middleware"
)

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteIdentityCookie(w, r)
	middleware.DeleteSessionCookie(w, r)

	http.Redirect(w, r, config.IDENTITY_LOGIN_PATH, http.StatusFound)
}
