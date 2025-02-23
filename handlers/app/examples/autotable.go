package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	. "previous/components"
	. "previous/handlers/app"

	// "previous/database"

	"previous/middleware"

	"net/http"
)

func AutoTableHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	func() Node {
		return AppLayout("Auto Table", *identity,
			P(Text("This codebase provides an API for generating filterable, sortable, and paginated datagrids such as the one shown below. You do not need to write a single line of JavaScript in order for this to work, as the \"interactivity\" is provided by HTMX.")),
			P(Text("Each interaction with an element of this table generates a dynamic SQL query.")),
			Br(),
			HxLoad("/app/examples/autotable-hx"),
		)
	}().Render(w)
}
