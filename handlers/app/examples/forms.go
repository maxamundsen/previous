package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/ui"

	"previous/middleware"

	"net/http"
)

type FormViewModel struct {
	display bool
	field1  string
	field2  string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	var viewModel FormViewModel

	if r.Method == http.MethodPost {
		r.ParseForm()

		viewModel.display = true
		viewModel.field1 = r.FormValue("field1")
		viewModel.field2 = r.FormValue("field2")
	}

	func() Node {
		return AppLayout("Form Submit Example", LAYOUT_SECTION_EXAMPLES, *identity, session,
			If(viewModel.display,
				Group{
					P(Text("You input:")),
					Ul(
						Li(InlineStyle("$me { color: $color(red-600);}"), Text(viewModel.field1)),
						Li(InlineStyle("$me { color: $color(red-600);}"), Text(viewModel.field2)),
					),
					Br(),
				},
			),
			Form(Method("post"), AutoComplete("off"),
				FormLabel(Text("Field 1")),
				FormInput(Type("text"), Name("field1")),

				FormLabel(Text("Field 2")),
				FormInput(Type("text"), Name("field2")),

				Br(),

				ButtonUI(Type("submit"), Text("Submit")),
			),
		)
	}().Render(w)
}
