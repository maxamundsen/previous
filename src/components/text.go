package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func PageLink(location string, display Node, newPage bool) Node {
	return A(Href(location), Class("underline text-neutral-600 hover:text-neutral-800"), display, If(newPage, Target("_blank")))
}
