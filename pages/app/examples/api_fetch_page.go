package examples

import (
	. "previous/components"
	. "previous/pages/app"

	. "maragu.dev/gomponents"

	"previous/middleware"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
func ApiFetchPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	func() Node {
		return AppLayout("API Fetch Example", *identity,
			PageLink("http://api.open-notify.org/astros.json", Text("http://api.open-notify.org/astros.json"), true),
			HxLoad("/app/examples/api-fetch-hx"),
		)
	}().Render(w)
}
