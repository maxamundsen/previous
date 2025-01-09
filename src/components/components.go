package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/basic"
	"previous/finance"
	"time"
)

// CONTAINERS
func Container(n ...Node) Node {
	return Div(Class("w-full px-3 mx-auto bs-sm:max-w-bs-sm bs-md:max-w-bs-md bs-lg:max-w-bs-lg bs-xl:max-w-bs-xl bs-xxl:max-w-bs-xxl"), Group(n))
}

func Card(header string, body ...Node) Node {
	return Div(Class("mt-5 p-10 bg-white border border-neutral-200 shadow"),
		H5(Class("mb-2 text-2xl font-bold text-neutral-900"), Text(header)),
		Group(body),
	)
}

// EMAIL
func ExampleEmailComponent(body string) Node {
	return EmailRoot(
		H1(Text("This email is automatically generated.")),
		P(Text(body)),
	)
}

// FORMATTERS
func FormatDateTime(utcTime time.Time) Node {
	return Text(TimeToString(utcTime))
}

func FormatDate(utcTime time.Time) Node {
	return Text(DateToString(utcTime))
}

func FormatMoney(m int64) Node {
	return Text(finance.Int64ToMoney(m))
}

func ToText(i interface{}) Node {
	return Text(ToString(i))
}


// FORMS
func FormInput(children ...Node) Node {
	return Input(Class("p-1 block w-full border-0 p-1.5 text-neutral-900 shadow-sm ring-1 ring-inset ring-neutral-300 placeholder:text-neutral-400 sm:text-sm/6"),
		Group(children),
	)
}

func FormSelect(children ...Node) Node {
	return Select(Class("p-3 bg-white block w-full border-0 p-1.5 text-neutral-900 shadow-sm ring-1 ring-inset ring-neutral-300 placeholder:text-neutral-400 sm:text-sm/6"),
		Group(children),
	)
}

func FormTextarea(children ...Node) Node {
	return Textarea(Class("block p-2.5 w-full text-sm text-neutral-900 bg-white shadow-sm border-0 ring-1 ring-inset ring-neutral-300"),
		Group(children),
	)
}

func FormLabel(children ...Node) Node {
	return Label(Class("text-neutral-900 text-sm"), Group(children))
}

// TEXT
func PageLink(location string, display Node, newPage bool) Node {
	return A(Href(location), Class("underline text-neutral-600 hover:text-neutral-800"), display, If(newPage, Target("_blank")))
}

// BUTTONS
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

// TABLES
func TableTW(c ...Node) Node {
	return Div(Class("flex flex-col"),
		Div(Class("-m-1.5 overflow-x-auto"),
			Div(Class("p-1.5 min-w-full inline-block align-middle"),
				Div(Class("overflow-hidden"),
					Table(Class("min-w-full divide-y divide-neutral-200 table-fixed"),
						Group(c),
					),
				),
			),
		),
	)
}

func TBodyTW(children ...Node) Node {
	return TBody(Class("divide-y divide-neutral-200"), Group(children))
}

func ThTW(children ...Node) Node {
	return Th(Class("px-6 py-3 text-start text-xs font-medium text-muted-foreground uppercase"), Group(children))
}

func TdTW(children ...Node) Node {
	return Td(Class("px-6 py-4 whitespace-nowrap text-sm font-medium text-muted-foreground"), Group(children))
}


