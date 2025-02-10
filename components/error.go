package components

import (
	"previous/.metagen/pageinfo"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ErrorPage(status int) Node {
	return RootLayout("An error has occurred.",
		Section(Class("bg-white"),
			Div(Class("py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6"),
				Div(Class("mx-auto max-w-screen-sm text-center"),
					H1(Class("mb-4 text-7xl tracking-tight lg:text-9xl text-neutral-950"), ToText(status)),
					P(Class("mb-4 text-lg text-neutral-500"), Text("We're sorry, an error has occured.")),
					A(Href(pageinfo.Root.Index.Url()), Class("inline-flex text-white bg-neutral-900 hover:bg-neutral-950 font-medium text-sm px-5 py-2.5 text-center my-4"), Text("Back to Homepage")),
				),
			),
		),
	)
}
