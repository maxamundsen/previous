package auth

import (
	"net/http"
	"saral/config"
	"saral/middleware"
)

// @Identity
// @Protected
// @Session
func LogoutController(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteIdentityCookie(w, r)
	middleware.DeleteSessionCookie(w, r)

	http.Redirect(w, r, config.IDENTITY_LOGIN_PATH, http.StatusFound)
}
