package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// TABLES
func TableTW(c ...Node) Node {
	return Div(Class("flex flex-col"),
		Div(Class("-m-1.5 overflow-x-auto"),
			Div(Class("p-1.5 min-w-full inline-block align-middle"),
				Div(Class("overflow-hidden"),
					Table(Class("min-w-full divide-y divide-gray-200 table-fixed"),
						Group(c),
					),
				),
			),
		),
	)
}

func TBodyTW(children ...Node) Node {
	return TBody(Class("divide-y divide-gray-200"), Group(children))
}

func ThTW(children ...Node) Node {
	return Th(Class("px-6 py-3 text-start text-xs font-medium text-muted-foreground uppercase"), Group(children))
}

func TdTW(children ...Node) Node {
	return Td(Class("px-6 py-4 whitespace-nowrap text-sm font-medium text-muted-foreground"), Group(children))
}
