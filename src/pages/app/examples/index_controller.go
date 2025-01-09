package examples

import (
	"net/http"
	"previous/auth"
	"previous/middleware"
	. "previous/pages/app"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// @Identity
// @Protected
// @CookieSession
func IndexController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	IndexView(*identity).Render(w)
}

func IndexView(identity auth.Identity) Node {
	return AppLayout("Example Index Page", identity,
		P(Text("This is an index page! Notice how the route is /app/examples without anything following?")),
	)
}