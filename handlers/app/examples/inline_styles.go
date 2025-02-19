package examples

import (
	"net/http"
	. "previous/components"
	. "previous/handlers/app"
	"previous/middleware"
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func InlineStylesHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	b, _ := strconv.ParseBool(r.URL.Query().Get("value"))

	func() Node {
		return AppLayout("Inline Styles", *identity,
			P(Text("This is another test page")),
			P(
				InlineStyle("me{font-size: var(--text-5xl);}"),

				// If `b` is true, make text green, else, make it red.
				IfElse(b,
					InlineStyle("me{color: var(--color-green-600);}"),
					InlineStyle("me{color: var(--color-red-600)}"),
				),

				Text("Inline styles can be applied conditionally. Click the buttons to change the color!"),
			),

			Form(Input(Type("hidden"), Value("true"), Name("value")), Button(Text("Make text green"))),
			Form(Input(Type("hidden"), Value("false"), Name("value")), Button(Text("Make text red"))),

			Br(),
			P((Text("* Note that these styles are determined server side."))),
		)
	}().Render(w)
}

func InlineStyleComponent() Node {
	return P(
		InlineStyle("me { color: blue; } @media md { me { color: red; padding: $(5); } }"),
		Text("Example!"),
	)
}
