package components

import (
	. "maragu.dev/gomponents"
)

// for some reason this isn't included in the gomponents package?
func MakeGroup(children ...Node) Node {
	return Group(children)
}

// also this...
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

