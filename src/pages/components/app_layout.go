package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AppLayout(title string, children ...Node) Node {
	return Root(title,
		AppNavbar(),
		Div(Class("container"),
			H3(Class("display-6"), Text(title)),
			Hr(),
			Group(children),
		),
		AppFooter(),
	)
}
