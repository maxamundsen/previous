package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ErrorPage(status int) Node {
	return Root("Error",
		Div(Class("container"), Style("margin-top: 100px"),
			H1(Class("text-danger"), Text("An error has ocurred.")),
			P(Text("HTTP Error: "), ToText(status)),
			Br(),
			A(Class("btn btn-secondary"), Href("/"), Text("Return to homepage")),
		),
	)
}
