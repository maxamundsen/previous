package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func PageLink(location string, display Node, newPage bool) Node {
	return A(Href(location), Class("underline text-red-600 hover:text-red-800"), display, If(newPage, Target("_blank")))
}
