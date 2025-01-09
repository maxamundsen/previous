package components

import (
	. "maragu.dev/gomponents"
)

func IfElse(condition bool, t Node, f Node) Node {
	if condition {
		return t
	} else {
		return f
	}
}

func IffElse(condition bool, t func() Node, f func() Node) Node {
	if condition {
		return t()
	} else {
		return f()
	}
}