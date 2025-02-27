package examples

import (
	. "previous/components"
	. "previous/handlers/app"

	. "maragu.dev/gomponents"

	"previous/middleware"

	"net/http"
)

func ApiFetchHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	AppLayout("API Fetch Example", *identity,
		PageLink("http://api.open-notify.org/astros.json", Text("http://api.open-notify.org/astros.json"), true),
		HxLoad("/app/examples/api-fetch-hx"),
	).Render(w)
}
