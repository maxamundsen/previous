package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Container(n ...Node) Node {
	return Div(Class("container mx-auto"), Group(n))
}