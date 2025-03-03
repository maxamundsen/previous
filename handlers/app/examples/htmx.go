package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/ui"

	"previous/auth"
	"previous/middleware"

	"net/http"
)

func HtmxHxHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	HtmxView(*identity, session).Render(w)
}

func HtmxView(identity auth.Identity, session map[string]interface{}) Node {
	return AppLayout("HTMX Example", LAYOUT_SECTION_EXAMPLES, identity, session,
		P(InlineStyle("$me{ margin-bottom: $5; }"), Text("Click the button to increase the counter. Open up the network tab in the browser developer tools to see how this works under the hood.")),
		CounterButton(0),

		Br(),
		Br(),

		Card(
			P(InlineStyle("$me{ margin-top: $5; margin-bottom: $5; }"), Text("HTMX is all about sending partial HTML snippets as the HTTP body, and swapping that response into the current DOM tree. This button is dynamically generated on the server, and HTMX automatically patches the DOM with the response. Unlike other AJAX methods, you specify the behavior entirely using HTML attributes, which follows LoB principles, and makes using HTMX very ergonomic.")),
			P(InlineStyle("$me { margin-bottom: $5; }"),
				Text("To learn more about HTMX, click "), PageLink("https://htmx.org", Text("here"), true), Text("."),
				Text(" This example only demonstrates functionality, and not utility. Check out the examples on the HTMX website for better uses of HTMX."),
			),
		),
	)
}
