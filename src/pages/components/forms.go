package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func FormInput(children ...Node) Node {
	return Input(Class("p-1 block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm/6"),
		Group(children),
	)
}

func FormSelect(children ...Node) Node {
	return Select(Class("p-3 bg-white block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm/6"),
		Group(children),
	)
}

func FormTextarea(children ...Node) Node {
	return Textarea(Class("block p-2.5 w-full text-sm text-gray-900 bg-white rounded-lg shadow-sm border-0 ring-1 ring-inset ring-gray-300"),
		Group(children),
	)
}

func FormLabel(children ...Node) Node {
	return Label(Class("text-gray-900 text-sm"), Group(children))
}
