package components

import (
	. "previous/basic"
	"previous/finance"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// DUMMY TEXT
const LOREM_IPSUM = ` Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam ultrices iaculis dui sed porttitor. Integer sed est fringilla, condimentum magna ac, sodales dui. Sed tempor pretium scelerisque. Vivamus pulvinar iaculis libero nec blandit. Mauris tempus velit in neque elementum, ac elementum diam feugiat. Aenean malesuada, nunc a interdum volutpat, diam est lacinia magna, nec fermentum massa lectus non urna. Cras vitae turpis finibus, porta est tincidunt, efficitur neque. Suspendisse suscipit a nulla mollis sodales. Nam vitae nulla vulputate, dictum purus eget, malesuada justo. Vestibulum ultricies eget neque ac volutpat. Mauris et molestie elit. Donec et suscipit urna. Duis in mi in ipsum faucibus finibus.`

// DIALOG / MODAL
func ModalActuator(id string, contents Node) Node {
	return Span(Attr("x-on:click", "$store."+id+" = true"),
		contents,
	)
}

func Modal(id string, title string, body Node, closeElements []Node) Node {
	return Div(
		Div(Attr("x-cloak", ""), Attr("x-show", "$store."+id), Attr("x-transition.opacity.duration.50ms", ""), Attr("x-trap.inert.noscroll", "$store."+id+""), Attr("x-on:keydown.esc.window", "$store."+id+" = false"), Attr("x-on:click.self", "$store."+id+" = false"), Class("fixed inset-0 z-30 flex items-end justify-center bg-black/20 p-4 pb-8 backdrop-blur-md sm:items-center lg:p-8"), Role("dialog"), Aria("modal", "true"), Aria("labelledby", "defaultModalTitle"),
			Div(Attr("x-show", "$store."+id+""), Attr("x-transition:enter", "transition ease-out duration-50 delay-20 motion-reduce:transition-opacity"), Attr("x-transition:enter-start", "opacity-0 scale-50"), Attr("x-transition:enter-end", "opacity-100 scale-100"), Class("flex max-w-lg flex-col gap-4 overflow-hidden rounded-radius border border-outline bg-surface text-on-surface dark:border-outline-dark dark:bg-surface-dark-alt dark:text-on-surface-dark"),
				Div(Class("flex items-center justify-between border-b border-outline bg-surface-alt/60 p-4"),
					H3(Class("font-semibold tracking-wide text-on-surface-strong"), Text(title)),
					Button(Class("cursor-pointer"), Attr("x-on:click", "$store."+id+" = false"), Aria("label", "close modal"),
						Icon(ICON_X_DIALOG_CLOSE, 24),
					),
				),
				Div(ID(id), Class("px-4 py-8"),
					body,
				),
				Div(Class("flex flex-col-reverse justify-between gap-2 border-t border-outline bg-surface-alt/60 p-4 sm:flex-row sm:items-center md:justify-end"),
					Map(closeElements, func(n Node) Node {
						return Div(Attr("x-on:click", "$store."+id+" = false"), n)
					}),
				),
			),
		),
	)
}

// CONTAINERS
func Container(n ...Node) Node {
	return Div(Class("w-full px-3 mx-auto bs-sm:max-w-bs-sm bs-md:max-w-bs-md bs-lg:max-w-bs-lg bs-xl:max-w-bs-xl bs-xxl:max-w-bs-xxl"), Group(n))
}

func Card(header string, body ...Node) Node {
	return Div(Class("mt-5 p-10 bg-white border border-neutral-200 shadow-sm"),
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
func FormatTime(utcTime time.Time) Node {
	return Text(TimeToTimeString(utcTime))
}

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
	return Input(Class("bg-white p-1 block w-full border-0 p-1.5 text-neutral-900 shadow-sm ring-1 ring-inset ring-neutral-300 placeholder:text-neutral-400 sm:text-sm/6"),
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
	return Button(Class("cursor-pointer group relative shadow inline-flex items-center overflow-hidden bg-neutral-600 px-8 py-1 text-white focus:outline-none focus:ring hover:bg-neutral-800 active:bg-neutral-800 text-sm"),
		Group(children),
	)
}

func ButtonRed(children ...Node) Node {
	return Button(Class("cursor-pointer group relative shadow inline-flex items-center overflow-hidden bg-red-600 px-8 py-1 text-white focus:outline-none focus:ring hover:bg-red-800 active:bg-red-800 text-sm"),
		Group(children),
	)
}

func ButtonBlue(children ...Node) Node {
	return Button(Class("cursor-pointer group relative shadow inline-flex items-center overflow-hidden bg-neutral-600 px-8 py-1 text-white focus:outline-none focus:ring hover:bg-neutral-800 active:bg-neutral-800 text-sm"),
		Group(children),
	)
}

// TABLES
func TableSearch(c ...Node) Node {
	return Input(Class("bg-white w-full pr-11 h-10 pl-3 py-2 bg-transparent placeholder:text-neutral-400 text-neutral-700 text-sm border border-neutral-200 transition duration-200 ease focus:outline-none focus:border-neutral-400 hover:border-neutral-400 shadow-sm focus:shadow-md"),
		Group(c),
	)
}

func TableSearchDropdown(c ...Node) Node {
	return Select(Class("bg-white w-full px-3 h-10 py-2 bg-transparent placeholder:text-neutral-400 text-neutral-700 text-sm border border-neutral-200 transition duration-200 ease focus:outline-none focus:border-neutral-400 hover:border-neutral-400 shadow-sm focus:shadow-md"),
		Group(c),
	)
}

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

// HTMX Helpers
func HxLoad(url string) Node {
	return Div(Attr("hx-get", url), Attr("hx-trigger", "load"),
		Spinner(),
	)
}

func Spinner() Node {
	// we could make this an "icon" but ... why?
	// also just return regular HTML because why not
	return Raw(`<div class="text-center"><div role="status"><svg aria-hidden="true" class="inline w-8 h-8 text-gray-200 animate-spin fill-gray-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/><path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/> </svg> <span class="sr-only">Loading...</span></div></div>`)
}

// NON Tailwind Components
func TestNoTw() Node {
	return Div(StyleAttr("color: red"),
		Text("hello"),
	)
}
