package components

import (
	. "previous/basic"
	"previous/finance"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// DUMMY TEXT
const LOREM_IPSUM = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam ultrices iaculis dui sed porttitor. Integer sed est fringilla, condimentum magna ac, sodales dui. Sed tempor pretium scelerisque. Vivamus pulvinar iaculis libero nec blandit. Mauris tempus velit in neque elementum, ac elementum diam feugiat. Aenean malesuada, nunc a interdum volutpat, diam est lacinia magna, nec fermentum massa lectus non urna. Cras vitae turpis finibus, porta est tincidunt, efficitur neque. Suspendisse suscipit a nulla mollis sodales. Nam vitae nulla vulputate, dictum purus eget, malesuada justo. Vestibulum ultricies eget neque ac volutpat. Mauris et molestie elit. Donec et suscipit urna. Duis in mi in ipsum faucibus finibus.`

// BADGES
func BadgeSuccess(children ...Node) Node {
	return Span(Class("w-fit inline-flex overflow-hidden rounded-sm border border-green-500 bg-white text-xs font-medium text-green-500 dark:border-green-500 dark:bg-neutral-950 dark:text-green-500"),
		Span(Class("px-2 py-1 bg-green-500/10 dark:bg-green-500/10"), Group(children)),
	)
}

func BadgeWarning(children ...Node) Node {
	return Span(Class("w-fit inline-flex overflow-hidden rounded-sm border border-amber-500 bg-white text-xs font-medium text-amber-500 dark:border-amber-500 dark:bg-neutral-950 dark:text-amber-500"),
		Span(Class("px-2 py-1 bg-amber-500/10 dark:bg-amber-500/10"), Group(children)),
	)
}

// DIALOG / MODAL
func ModalActuator(id string, contents Node) Node {
	return Span(Attr("x-on:click", "$store."+id+" = true"),
		contents,
	)
}

func Modal(id string, title string, body Node, closeElements []Node) Node {
	return Div(
		Div(Attr("x-cloak", ""), Attr("x-show", "$store."+id), Attr("x-transition.opacity.duration.50ms", ""), Attr("x-trap.inert.noscroll", "$store."+id+""), Attr("x-on:keydown.esc.window", "$store."+id+" = false"), Attr("x-on:click.self", "$store."+id+" = false"), Class("fixed inset-0 z-30 flex items-end justify-center bg-black/30 p-4 pb-8 sm:items-center lg:p-8"), Role("dialog"), Aria("modal", "true"), Aria("labelledby", "defaultModalTitle"),
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
	return Div(
		InlineStyle(`
			me > input {
				background-color: var(--color-white);
				padding: $(2);
				display: block;
				width: 100%;
				border: 0;
				color: var(--color-neutral-900);
				box-shadow: var(--shadow-sm);
			}

			@media sm {
				me >input {
					font-size: var(--text-sm);
				}
			}
		`),
		Input(Group(children)),
	)
}

func FormSelect(children ...Node) Node {
	return Div(
		InlineStyle(`
			me > select {
				padding: $(3);
				background-color: var(--color-white);
				display: block;
				width: 100%;
				border: 0;
				color: var(--color-neutral-900);
				box-shadow: var(--shadow-sm);
			}
			@media sm {
				me > select{
					font-size: var(--text-sm);
				}
			}
		`),
		Select(Group(children)),
	)
}

func FormTextarea(children ...Node) Node {
	return Div(
		InlineStyle(`
			me > textarea {
				display: block;
				padding: $(3);
				width: 100%;
				font-size: var(--text-sm);
				color: var(--color-neutral-900);
				background-color: var(--color-white);
				box-shadow: var(--shadow-sm);
			}
		`),
		Textarea(
			Group(children),
		),
	)
}

func FormLabel(children ...Node) Node {
	return Label(Class("text-neutral-900 text-sm"), Group(children))
}

// TEXT
func PageLink(location string, display Node, newPage bool) Node {
	return A(Href(location), InlineStyle("me{text-decoration: underline; color: var(--color-blue-600);} this:hover{color: var(--color-blue-800);}"), display, If(newPage, Target("_blank")))
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
	return Button(Class("cursor-pointer group relative shadow inline-flex items-center overflow-hidden bg-blue-600 px-8 py-1 text-white focus:outline-none focus:ring hover:bg-blue-800 active:bg-neutral-800 text-sm"),
		Group(children),
	)
}

// TABLES
func TableSearchDropdown(c ...Node) Node {
	return Select(Class("bg-white w-full px-3 h-10 py-2 bg-transparent placeholder:text-neutral-400 text-neutral-700 text-sm border border-neutral-200 transition duration-200 ease focus:outline-none focus:border-neutral-400 hover:border-neutral-400 shadow-sm focus:shadow-md"),
		Group(c),
	)
}

func TableTW(c ...Node) Node {
	return Div(Class("flex flex-col"),
		Div(InlineStyle("me {margin: $(2); overflow: x-auto;}"),
			Div(InlineStyle("me { padding: $(2); min-width: 100%; display: inline-block; vertical-align: middle; }"),
				Div(InlineStyle("me { overflow: hidden; }"),
					Table(
						InlineStyle(`
							me {
								min-width: 100%;
								table-layout: fixed;
							}

							me > :not(:last-child) {
								border-top-width: 0px;
								border-bottom-width: 1px;
							}
						`),
						Group(c),
					),
				),
			),
		),
	)
}

func TBodyTW(children ...Node) Node {
	return TBody(
		InlineStyle(`
			me > :not(:last-child) {
				border-top-width: 0px;
				border-bottom-width: 1px;
				border-color: var(--color-neutral-500);
			}
		`),
		Group(children),
	)
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
		Loader(),
	)
}

func Loader() Node {
	return Span(
		InlineStyle(`
		me {
			width: 48px;
			height: 48px;
			border-radius: 50%;
			display: inline-block;
			border-top: 4px solid #FFF;
			border-right: 4px solid transparent;
			box-sizing: border-box;
			animation: rotation 1s linear infinite;
		}
		me::after {
			content: '';
			box-sizing: border-box;
			position: absolute;
			left: 0;
			top: 0;
			width: 48px;
			height: 48px;
			border-radius: 50%;
			border-bottom: 4px solid #FF3D00;
			border-left: 4px solid transparent;
		}
		@keyframes rotation {
			0% {
				transform: rotate(0deg);
			}
			100% {
				transform: rotate(360deg);
			}
		}
		`),
	)
}
