package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
	"previous/middleware"
	. "previous/pages/app"
	. "previous/components"
)

// @Identity
// @Protected
// @CookieSession
func InlineStylesPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	b := true // change me to true

	func() Node {
		return AppLayout("Inline Styles", *identity,
			P(Text("This is another test page")),
			P(
				InlineStyle("$this{font-size: var(--text-5xl);}"),
				IfElse(b, InlineStyle("$this{color: var(--color-green-600);}"), InlineStyle("$this{color: var(--color-red-600)}")),
				Text("You can change and append styles based on conditions by chaining InlineStyle calls together."),
			),
			Br(),
			P(Text("* Note that these styles are determined server side.")),
		)
	}().Render(w)
}
