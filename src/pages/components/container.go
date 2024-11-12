package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Container(n ...Node) Node {
	return Div(Class("w-full px-3 mx-auto bs-sm:max-w-bs-sm bs-md:max-w-bs-md bs-lg:max-w-bs-lg bs-xl:max-w-bs-xl bs-xxl:max-w-bs-xxl"), Group(n))
}
