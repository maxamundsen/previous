package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
	. "previous/ui"
	"previous/middleware"
)

func QuillHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	var input string

	if r.Method == http.MethodPost {
		r.ParseForm()
	}

	AppLayout("WYSIWYG Editor - Quill.js", LAYOUT_SECTION_EXAMPLES, *identity, session,
		Form(Action(""), Method("POST"),
			Quill(),
			Br(),
			ButtonUI(Type("submit"), Text("Render")),
		),
		If(input != "",
			Group{
				Br(),
				Card(
					Prose(
						SafeRaw(""),
					),
				),
			},
		),
	).Render(w)
}
