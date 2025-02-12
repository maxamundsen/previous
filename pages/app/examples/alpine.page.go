package examples

import (
	. "previous/components"
	. "previous/pages/app"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/middleware"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
// @Static
func AlpinePage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	data := make(AlpineStore)
	data["some_value"] = "10"

	func() Node {
		return AppLayout("Alpine Example", *identity,
			// Map store values from Go to Javascript
			data.Init(),

			Div(Attr("x-data", "{ count: 0 }"),
				P(Class("mb-5"),
					Text("Click the button to increase the counter. This interaction relies on client-side scripting."),
					Text(" Alpine.js allows for simple DOM manipulation using HTML attributes."),
					Text(" Although Previous isn't built using JavaScript, it is still sometimes necessary for features such as clickable dropdown menues, or modal dialogs."),
				),
				ButtonGray(Attr("x-text", `"Counter: " + count`), Attr("x-on:click", "count+=1")),
				P(Class("mt-5"),
					Text("To learn more about Alpine.js, click "), PageLink("https://alpinejs.dev", Text("here"), true), Text("."),
				),
			),
			Div(
				FormInput(Attr("x-bind", "$store.some_value")),
				P(Attr("x-html", "$store.some_value")),
				H1(Attr("x-html", "$store.some_value")),
			),
		)
	}().Render(w)
}
