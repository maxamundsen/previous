package public

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	IndexView().Render(w)
}

func IndexView() Node {
	return PublicLayout("Index Page",
		Br(),
		Br(),
		Div(Class("text-center"),
			H1(Text("This is a demo website.")),
			Br(),
			Br(),
			A(Href("/auth/login"), Button(Class("btn btn-primary"), Text("Sign in"))),
		),
	)
}
