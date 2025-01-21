package examples

import (
	. "previous/components"
	. "previous/pages/app"

	. "maragu.dev/gomponents"

	"previous/auth"
	"previous/middleware"

	"net/http"
)

// Struct for deserializing json from 3rd party API
// http://api.open-notify.org/astros.json
type astroModel struct {
	Message string   `json:"message"`
	People  []person `json:"people"`
	Number  int      `json:"number"`
}

type person struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

// @Identity
// @Protected
// @CookieSession
func ApiFetchController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	ApiFetchView(identity).Render(w)
}

func ApiFetchView(identity *auth.Identity) Node {
	return AppLayout("API Fetch Example", *identity,
		PageLink("http://api.open-notify.org/astros.json", Text("http://api.open-notify.org/astros.json"), true),
		HxLoad("/app/examples/api-fetch-hx"),
	)
}
