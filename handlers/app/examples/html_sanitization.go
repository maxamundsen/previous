package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
	. "previous/ui"
	"previous/middleware"
)

func HtmlSanitizationHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	input := `<h1>This is user input HTML</h1>
<a href="/app/dashboard">Go to dashboard page!</a>

<script>window.alert("XSS is protected against! This doesn't work.")</script>
`
	if r.Method == http.MethodPost {
		r.ParseForm()
		input = r.FormValue("html_content")
	}

	AppLayout("HTML Sanitization", LAYOUT_SECTION_EXAMPLES, *identity, session,
		Form(Action(""), Method("POST"),
			FormTextarea(Name("html_content"), Placeholder("Type HTML input here:"), Text(input), StyleAttr("height: 400px;")),
			Br(),
			ButtonUI(Type("submit"), Text("Render")),
		),
		If(input != "",
			Group{
				Br(),
				Card(
					Prose(
						SafeRaw(input),
					),
				),
			},
		),
		InlineScript("hljs.highlightAll();"),
	).Render(w)
}
