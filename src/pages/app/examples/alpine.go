package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "saral/pages/components"

	"saral/middleware"
	"saral/models"

	"net/http"
)

func AlpineController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	AlpineView(*identity).Render(w)
}

func AlpineView(identity models.Identity) Node {
	return AppLayout("Alpine Example", identity,
		Div(Attr("x-data", "{ count: 0 }"),
			P(Class("mb-5"),
				Text("Click the button to increase the counter. This interaction relies on client-side scripting."),
				Text(" Alpine.js allows for simple DOM manipulation using HTML attributes."),
				Text(" Although Saral isn't built using JavaScript, it is still sometimes necessary for features such as clickable dropdown menues, or modal dialogs."),
			),
			ButtonGray(Attr("x-text", `"Counter: " + count`), Attr("x-on:click", "count+=1")),
			P(Class("mt-5"),
				Text("To learn more about Alpine.js, click "), PageLink("https://alpinejs.dev", Text("here"), true), Text("."),
			),
		),
	)
}
