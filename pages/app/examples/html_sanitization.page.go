package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
	"previous/middleware"
	. "previous/components"
	. "previous/pages/app"
)

// @Identity
// @Protected
// @CookieSession
func HtmlSanitizationPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	input := `<h1>This is user input HTML</h1>
<a href="/app/dashboard">Go to dashboard page!</a>

<script>window.alert("XSS is protected against! This doesn't work.")</script>
`
    if r.Method == http.MethodPost {
        r.ParseForm()
        input = r.FormValue("html_content")
    }

	func() Node {
		return AppLayout("HTML Sanitization", *identity,
			Form(Action(""), Method("POST"),
				FormTextarea(Name("html_content"), Placeholder("Type HTML input here:"), Text(input), StyleAttr("height: 400px;")),
				Br(),
				ButtonGray(Type("submit"), Text("Render")),
			),
			If(input != "",
				Group{
					Br(),
					Card("",
						Div(StyleReset(),
							SafeRaw(input),
						),
					),
				},
			),
			Script(Raw("hljs.highlightAll();")),
		)
	}().Render(w)
}
