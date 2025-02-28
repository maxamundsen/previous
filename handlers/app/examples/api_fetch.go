package examples

import (
	. "previous/components"

	. "maragu.dev/gomponents"

	"previous/middleware"

	"net/http"
)

func ApiFetchHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	AppLayout("API Fetch Example", LAYOUT_SECTION_EXAMPLES, *identity, session,
		PageLink("http://api.open-notify.org/astros.json", Text("http://api.open-notify.org/astros.json"), true),
		HxLoad("/app/examples/api-fetch-hx"),
	).Render(w)
}
