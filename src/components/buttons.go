package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ButtonGray(children ...Node) Node {
	return Button(Class("group relative shadow inline-flex items-center overflow-hidden bg-neutral-600 px-8 py-1 text-white focus:outline-none focus:ring hover:bg-neutral-800 active:bg-neutral-800 text-sm"),
		Group(children),
	)
}

func ButtonRed(children ...Node) Node {
	return Button(Class("group relative shadow inline-flex items-center overflow-hidden bg-red-600 px-8 py-1 text-white focus:outline-none focus:ring hover:bg-red-800 active:bg-red-800 text-sm"),
		Group(children),
	)
}

func ButtonBlue(children ...Node) Node {
	return Button(Class("group relative shadow inline-flex items-center overflow-hidden bg-neutral-600 px-8 py-1 text-white focus:outline-none focus:ring hover:bg-neutral-800 active:bg-neutral-800 text-sm"),
		Group(children),
	)
}
