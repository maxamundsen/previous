package examples

import (
	. "previous/components"
	"previous/.metagen/pageinfo"
	. "previous/pages/app"

	. "maragu.dev/gomponents"

	"previous/middleware"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
// @Static
func ApiFetchPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	func() Node {
		return AppLayout("API Fetch Example", *identity,
			PageLink("http://api.open-notify.org/astros.json", Text("http://api.open-notify.org/astros.json"), true),
			HxLoad(pageinfo.Root.App.Examples.Api_fetch_hx.Url()),
		)
	}().Render(w)
}
