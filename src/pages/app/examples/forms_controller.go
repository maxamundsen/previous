package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/components"
	. "previous/pages/app"

	"previous/auth"
	"previous/middleware"

	"net/http"
)

type FormViewModel struct {
	display bool
	field1 string
	field2 string
}

// @Identity
// @Protected
// @CookieSession
func FormController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	var viewModel FormViewModel

	if r.Method == http.MethodPost {
		r.ParseForm()

		viewModel.display = true
		viewModel.field1 = r.FormValue("field1")
		viewModel.field2 = r.FormValue("field2")
	}

	FormView(&viewModel, *identity).Render(w)
}

func FormView(viewModel *FormViewModel, identity auth.Identity) Node {
	return AppLayout("Form Submit Example", identity,
		If(viewModel.display,
			Group{
				P(Text("You input:")),
				Ul(
					Li(Class("text-red-400"), Text(viewModel.field1)),
					Li(Class("text-red-400"), Text(viewModel.field2)),
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

			ButtonGray(Type("submit"), Text("Submit")),
		),
	)
}