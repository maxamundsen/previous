package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Card(header string, body ...Node) Node {
	return Div(Class("mt-5 p-10 bg-white border border-neutral-200 shadow"),
		H5(Class("mb-2 text-2xl font-bold text-neutral-900"), Text(header)),
		Group(body),
	)
}
