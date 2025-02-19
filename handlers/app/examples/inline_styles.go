package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
	"previous/middleware"
	. "previous/handlers/app"
	. "previous/components"
)

func InlineStylesHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	b := true // change me to true

	func() Node {
		return AppLayout("Inline Styles", *identity,
			P(Text("This is another test page")),
			P(
				InlineStyle("me{font-size: var(--text-5xl);}"),
				IfElse(b, InlineStyle("me{color: var(--color-green-600);}"), InlineStyle("me{color: var(--color-red-600)}")),
				Text("You can change and append styles based on conditions by chaining InlineStyle calls together."),
			),
			Br(),
			P(Text("* Note that these styles are determined server side.")),
		)
	}().Render(w)
}

func InlineStyleComponent() Node {
	return P(
		InlineStyle("me { color: blue; } @media md { me { color: red; padding: $(5); } }"),
		Text("Example!"),
	)
}