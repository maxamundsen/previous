package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func InlineStyle(input string) Node {
	return StyleEl(Raw(input))
}