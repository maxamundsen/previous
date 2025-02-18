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

func AutoTablePage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	AutoTableView(*identity).Render(w)
}

func AutoTableView(identity auth.Identity) Node {
	return AppLayout("Auto Table", identity,
		P(Text("This codebase provides an API for generating filterable, sortable, and paginated datagrids such as the one shown below. You do not need to write a single line of JavaScript in order for this to work, as the \"interactivity\" is provided by HTMX.")),
		P(Text("Each interaction with an element of this table generates a dynamic SQL query.")),
		Br(),
		HxLoad("/app/examples/autotable-hx"),
	)
}