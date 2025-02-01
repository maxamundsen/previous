package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	. "previous/components"
	. "previous/pages/app"

	// "previous/repository"

	"previous/auth"
	"previous/middleware"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
func ServerSideDataGridPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	PaginationView(*identity).Render(w)
}

func PaginationView(identity auth.Identity) Node {
	return AppLayout("Server Side Data Grid", identity,
		P(Text("Previous provides an API for generating filterable, sortable, and paginated datagrids such as the one shown below. You do not need to write a single line of JavaScript in order for this to work, as the \"interactivity\" is provided by HTMX.")),
		Br(),
		HxLoad("/app/examples/orders-hx"),
	)
}