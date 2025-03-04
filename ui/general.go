package ui

import (
	. "previous/basic"
	"previous/finance"
	"time"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

// DUMMY TEXT
const LOREM_IPSUM = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam ultrices iaculis dui sed porttitor. Integer sed est fringilla, condimentum magna ac, sodales dui. Sed tempor pretium scelerisque. Vivamus pulvinar iaculis libero nec blandit. Mauris tempus velit in neque elementum, ac elementum diam feugiat. Aenean malesuada, nunc a interdum volutpat, diam est lacinia magna, nec fermentum massa lectus non urna. Cras vitae turpis finibus, porta est tincidunt, efficitur neque. Suspendisse suscipit a nulla mollis sodales. Nam vitae nulla vulputate, dictum purus eget, malesuada justo. Vestibulum ultricies eget neque ac volutpat. Mauris et molestie elit. Donec et suscipit urna. Duis in mi in ipsum faucibus finibus.`

// Splitters / dividers / section splits
func Divider() Node {
	return Hr(
		InlineStyle("$me { color: $color(neutral-200); margin-bottom: $3; margin-top: $1; }"),
	)
}

// By default, HTML element styles are reset by tailwind's preflight css
// (we use their styles even though we aren't using tailwind).
// This element reverts all child elements back to the _browser_ defaults,
// and applies additional styles to make user-input HTML look nicer.
//
// Useful for markdown content, blogs, etc.
func Prose(children ...Node) Node {
	return Div(
		InlineStyle(`
			$me * {
				all: revert;
				font-family: var(--font-sans);
			}
		`),
		Group(children),
	)
}

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
	return Span(
		InlineScriptf(`
			let act = me();
			let dialog = me("#%s_dialog");
			act.on("click", () => { dialog.showModal(); });
		`, id),
		contents,
	)
}

func Modal(id string, header Node, body Node, closeElements []Node) Node {
	return Dialog(
		ID(id+"_dialog"),
		InlineStyle(`
			$me {
				top: 50%;
				left: 50%;
				translate: -50% -50%;
				z-index: 10;
				background: $color(white);
				box-shadow: var(--shadow-md);
				border-bottom: 3px solid $color(neutral-600);
				border-radius: var(--radius-xs);
				font-size: var(--text-sm);
			}

			$me::backdrop {
				background: $color(black/30);
			}
		`),

		// Header
		Div(
			InlineStyle(`
				$me {
					display: flex;
					align-items: center;
					justify-content: space-between;
					padding: $6 $8 $1 $8;
					font-size: var(--text-xl);
					font-weight: var(--font-weight-bold);
				}
			`),
			header,

			Div(
				Class("modal-close-btn"),
				InlineStyle("$me { cursor: pointer; color: $color(gray-500); }"),
				Icon(ICON_X_DIALOG_CLOSE, 24),
			),
		),

		//Modal contents
		Div(
			ID(id), // we use the passed id here so that swapping the content is easier
			InlineStyle(`
				$me {
					padding: $3 $8 $8 $8;
					color: $color(neutral-700);
				}
			`),
			body,
		),

		// footer
		If(len(closeElements) > 0,
			Div(
				InlineStyle(`
					$me {
						padding: $3 $8 $3 $8;
						display: flex;
						flex-direction: row-reverse;
						gap: $2;
					}
				`),

				Map(closeElements, func(el Node) Node {
					return Div(Class("modal-close-btn"),
						el,
					)
				}),
			),
		),

		InlineScript(`
			let dialog = me();
			let close_buttons = any(".modal-close-btn", me())

			dialog.on("click", (ev) => {
				if (ev.target === dialog) {
					dialog.close();
				}
			});

			close_buttons.on("click", () => { dialog.close(); });
		`),
	)
}

// CONTAINERS
func Flex(n ...Node) Node {
	return Div(
		InlineStyle("$me { display: flex; align-items: center; flex-direction: row; gap: $3; }"),
		Group(n),
	)
}

func CardNoPadding(body ...Node) Node {
	return Div(
		InlineStyle(`
			$me {
				background-color: $color(white);
				box-shadow: var(--shadow-xs);
			}
		`),
		Group(body),
	)
}
func Card(body ...Node) Node {
	return Div(
		InlineStyle(`
			$me {
				background-color: $color(white);
				padding: $5;
				box-shadow: var(--shadow-xs);
			}
		`),
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

// TEXT
func PageLink(location string, display Node, newPage bool) Node {
	return A(
		Href(location),
		InlineStyle("$me{text-decoration: underline; color: $color(blue-600);} $me:hover{color: $color(blue-800);}"),
		display,
		If(newPage, Target("_blank")),
	)
}

// TABLES
func TdMoney(amt int64) Node {
	return TdRight(
		InlineStyle("$me { display: flex; justify-content: space-between;}"),
		Div(Text("$")),
		B(FormatMoney(int64(amt))),
	)
}

// Row Item Helpers
func TdLeft(c ...Node) Node {
	return Td(InlineStyle("$me { text-align: left; }"), Group(c))
}

func TdRight(c ...Node) Node {
	return Td(InlineStyle("$me { text-align: right; }"), Group(c))
}

// HTMX Helpers
func HxLoad(url string) Node {
	return Div(hx.Get(url), hx.Trigger("load"),
		Loader(),
	)
}

func Loader() Node {
	return Div(
		InlineStyle(`
		$me {
		    width: 48px;
		    height: 48px;
		    border: 5px solid #FFF;
		    border-bottom-color: $color(neutral-800);
		    border-radius: 50%;
		    display: inline-block;
		    box-sizing: border-box;
		    animation: rotation 1s linear infinite;
		    margin: 0 auto;
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
