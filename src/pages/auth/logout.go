package auth

import (
	"net/http"
	"webdawgengine/config"
	"webdawgengine/middleware"
)

func LogoutController(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteIdentityCookie(w, r)
	middleware.DeleteSessionCookie(w, r)

	http.Redirect(w, r, config.IDENTITY_LOGIN_PATH, http.StatusFound)
}
