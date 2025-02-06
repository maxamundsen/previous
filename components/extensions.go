// This file contains a few extensions to the "gomponents" library.
// I deliberately keep these extensions separated from the library to indicate

package components

import (
	. "maragu.dev/gomponents"
)

func MapMapWithKey[T comparable, U comparable](m map[T]U, cb func(T, U) Node) Group {
	var nodes []Node

	for k, _ := range m {
		nodes = append(nodes, cb(k, m[k]))
	}

	return nodes
}

func MapMap[T comparable, U comparable](m map[T]U, cb func(U) Node) Group {
	var nodes []Node

	for k, _ := range m {
		nodes = append(nodes, cb(m[k]))
	}

	return nodes
}

func MapWithIndex[T any](collection []T, callback func(index int, item T) Node) Group {
	nodes := make([]Node, 0, len(collection))
	for index, item := range collection {
		nodes = append(nodes, callback(index, item))
	}
	return nodes
}

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

func Template(children ...Node) Node {
	return El("template", children...)
}
