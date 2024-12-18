package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ErrorPage(status int) Node {
	return Root("An error has occurred.",
		Section(Class("bg-white"),
			Div(Class("py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6"),
				Div(Class("mx-auto max-w-screen-sm text-center"),
					H1(Class("mb-4 text-7xl tracking-tight font-serif lg:text-9xl text-blue-950"), ToText(status)),
					P(Class("mb-4 text-lg text-gray-500"), Text("We're sorry, an error has occured.")),
					A(Href("/"), Class("inline-flex text-white bg-blue-900 hover:bg-blue-950 font-medium text-sm px-5 py-2.5 text-center my-4"), Text("Back to Homepage")),
				),
			),
		),
	)
}
