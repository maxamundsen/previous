package middleware

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"webdawgengine/config"
	"webdawgengine/database"
	"webdawgengine/models"
)

type identityKey struct{}

func NewIdentity(userid int, securityStamp string, authenticated bool, rememberMe bool) *models.Identity {
	expirationDuration := time.Duration(time.Hour * 24 * time.Duration(config.IDENTITY_COOKIE_EXPIRY_DAYS))
	expiration := time.Now().Add(expirationDuration)

	return &models.Identity{
		UserId:        userid,
		SecurityStamp: securityStamp,
		Authenticated: authenticated,
		RememberMe:    rememberMe,
		Expiration:    expiration,
	}
}

func LoadIdentity(h http.HandlerFunc, requireAuth bool) http.HandlerFunc {
	loginPath := config.IDENTITY_LOGIN_PATH
	logoutPath := config.IDENTITY_LOGOUT_PATH
	defaultPath := config.IDENTITY_DEFAULT_PATH
	redirect := config.IDENTITY_AUTH_REDIRECT

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var identity *models.Identity

		token := r.Header.Get("Authorization")

		// if bearer token present, use token auth, else use cookies
		if token != "" {
			splitToken := strings.Split(token, "Bearer ")

			if len(splitToken) >= 2 {
				token = splitToken[1]
				identity, _ = DecryptIdentity(token)
			}

			if identity == nil {
				blankIdentity := &models.Identity{Authenticated: false}

				if requireAuth {
					http.Error(w, "Error: Unauthorized", http.StatusUnauthorized)
					return
				}

				ctx := context.WithValue(r.Context(), identityKey{}, blankIdentity)
				h.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			securityStamp, _ := database.FetchUserSecurityStamp(identity.UserId)

			securityCheckFailed := securityStamp != identity.SecurityStamp
			notAuthenticated := requireAuth && !identity.Authenticated
			identityExpired := identity.Expiration.Before(time.Now())

			if securityCheckFailed || notAuthenticated || identityExpired {
				http.Error(w, "Error: Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), identityKey{}, identity)
			h.ServeHTTP(w, r.WithContext(ctx))
		} else {
			identityCookie, err := r.Cookie(config.IDENTITY_COOKIE_NAME)
			if err == nil {
				identity, _ = DecryptIdentity(identityCookie.Value)
			}

			if identity == nil {
				blankIdentity := &models.Identity{Authenticated: false}

				if requireAuth {
					if redirect && loginPath != r.URL.Path && logoutPath != r.URL.Path {
						redirectPath := loginPath + "?redirect=" + url.QueryEscape(r.URL.String())
						http.Redirect(w, r, redirectPath, http.StatusFound)

						return
					} else if !redirect && loginPath != r.URL.Path {
						http.Error(w, "Error: Unauthorized", http.StatusUnauthorized)
						return
					}
				}

				ctx := context.WithValue(r.Context(), identityKey{}, blankIdentity)
				h.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			securityStamp, _ := database.FetchUserSecurityStamp(identity.UserId)

			securityCheckFailed := securityStamp != identity.SecurityStamp
			notAuthenticated := requireAuth && !identity.Authenticated
			identityExpired := identity.Expiration.Before(time.Now())

			if securityCheckFailed || notAuthenticated || identityExpired {
				DeleteIdentityCookie(w, r)
				http.Redirect(w, r, loginPath, http.StatusFound)
				return
			}

			if redirect && loginPath == r.URL.Path {
				http.Redirect(w, r, defaultPath, http.StatusFound)
				return
			}

			ctx := context.WithValue(r.Context(), identityKey{}, identity)
			h.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

func GetIdentity(r *http.Request) *models.Identity {
	identity := r.Context().Value(identityKey{}).(*models.Identity)
	return identity
}

func PutIdentityCookie(w http.ResponseWriter, r *http.Request, identity *models.Identity) {
	cookies := r.Cookies()

	// calculate total bytes used by other cookies
	var totalBytes int
	for _, cookie := range cookies {
		if cookie.Name == config.IDENTITY_COOKIE_NAME {
			continue
		} else {
			totalBytes += len(cookie.Value)
		}
	}

	cookieString, err := EncryptIdentity(identity)
	if err != nil {
		return
	}

	length := len(cookieString) + 8 // 8 additional bytes coming from somewhere ¯\_(ツ)_/¯

	if length+totalBytes > 4096 {
		log.Println("Attempt to generate cookie exceeding 4096 bytes")
		return
	}

	httpCookie := &http.Cookie{
		Name:     config.IDENTITY_COOKIE_NAME,
		Value:    cookieString,
		HttpOnly: true,
		Secure:   r.URL.Scheme == "https",
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	// if no expiry is set, cookie defaults to clear after browser closes

	if identity.RememberMe {
		httpCookie.Expires = identity.Expiration
	}

	http.SetCookie(w, httpCookie)
}

func DeleteIdentityCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    config.IDENTITY_COOKIE_NAME,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
		Path:    "/",
	})
}
