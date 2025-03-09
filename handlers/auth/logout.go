package auth

import (
	"net/http"
	"previous/constants"
	"previous/middleware"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	middleware.DeleteIdentityCookie(w, r)
	middleware.DeleteSessionCookie(w, r)

	http.Redirect(w, r, constants.IDENTITY_LOGIN_PATH, http.StatusFound)
}
