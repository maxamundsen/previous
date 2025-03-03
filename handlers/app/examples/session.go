package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	. "previous/ui"

	// "previous/database"

	"previous/middleware"

	"net/http"
)

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	const (
		SESSION_KEY = "EXAMPLE_ITEM"
	)

	if r.Method == http.MethodPost {
		r.ParseForm()
		token := r.FormValue("session_item")
		session[SESSION_KEY] = token
		middleware.PutSessionCookie(w, r, session)
	}

	_, ok := session[SESSION_KEY]

	AppLayout("Cookie Sessions", LAYOUT_SECTION_EXAMPLES, *identity, session,
		Card(
			P(
				Text("Sometimes, you want to store data between pages that is specific to the users login session. "),
				Text("Storing this data in the database since it is not long-lived, and will reset when the user logs out."),
			),

			P(
				Text("To solve this problem, we have \"cookie sessions\" which allow you to store arbitrary data in a cookie, which is sent back to the server each request. "),
				Text(""),
			),
		),

		Br(),
		Br(),

		Form(
			Method("POST"),
			Action(r.URL.Path),
			FormInput(
				Type("text"),
				Name("session_item"),
				Placeholder("Enter value to store..."),
			),
			Br(),

			ButtonUI(Type("submit"), Text("Push Session Cookie")),
		),

		Br(),
		Br(),

		If(ok && session[SESSION_KEY] != "",
			Card(
				P(
					InlineStyle("$me { color: $color(neutral-600); }"),
					Em(Text("Leave this page and come back, and the value will still be there. This value will be cleared on logout.")),
				),
				ToText(session[SESSION_KEY]),
			),
		),

	).Render(w)
}
