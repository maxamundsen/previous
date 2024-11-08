package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Card(header string, body ...Node) Node {
	return Div(Class("mt-5 p-10 bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700"),
		H5(Class("mb-2 text-2xl font-bold text-gray-900 dark:text-white"), Text(header)),
		Group(body),
	)
}