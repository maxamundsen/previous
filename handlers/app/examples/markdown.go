package examples

import (
	"net/http"
	. "previous/ui"
	"previous/middleware"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func MarkdownHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	input := `# Markdown Test

**This is an example of Markdown Rendering.**
Press the ` + "`" + `Render` + "`" + ` button to render this body to HTML.

### TODO List

- Build things
- Ship them
- Win

## Code Example
_Note that this page also uses highlight.js for highlighting code._

` + "```" + `go
func Sum(a, b int) int {
	return a + b
}
` + "```" + `

[XSS LINK](javascript:alert('User input protected against XSS. This doesn't work.'))

<script>console.log("XSS is protected against. This message will NOT be evaluated.")</script>
`

	if r.Method == http.MethodPost {
		r.ParseForm()
		input = r.FormValue("md_content")
	}

	func() Node {
		return AppLayout("Markdown Rendering", LAYOUT_SECTION_EXAMPLES, *identity, session,
			Form(Action(""), Method("POST"),
				FormTextarea(Name("md_content"), Placeholder("Type markdown input here:"), Text(input), StyleAttr("height: 400px;")),
				Br(),
				ButtonUI(Type("submit"), Text("Render")),
			),
			If(input != "",
				Group{
					Br(),
					Card(
						Prose(
							Markdown(input),
						),
					),
				},
			),
			InlineScript("hljs.highlightAll();"),
		)
	}().Render(w)
}
