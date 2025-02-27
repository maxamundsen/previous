package examples

import (
	"net/http"
	. "previous/components"
	. "previous/handlers/app"
	"previous/middleware"
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func InlineStylesHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	b, _ := strconv.ParseBool(r.URL.Query().Get("value"))

	AppLayout("Inline Styles", *identity,
		P(Text("This is another test page")),
		P(
			InlineStyle("$me{font-size: var(--text-5xl);}"),

			// If `b` is true, make text green, else, make it red.
			IfElse(b,
				InlineStyle("$me{color: $color(green-600);}"),
				InlineStyle("$me{color: $color(red-600);}"),
			),

			Text("Inline styles can be applied conditionally. Click the buttons to change the color!"),
		),

		Form(Input(Type("hidden"), Value("true"), Name("value")), ButtonUI(Text("Make text green"))),
		Br(),
		Form(Input(Type("hidden"), Value("false"), Name("value")), ButtonUI(Text("Make text red"))),

		Br(),
		P((Text("* Note that these styles are determined server side."))),
		Br(),

		MacroDescription("$me", "CSS selector for the current element you are inside."),
		MacroDescription("$color(<color>/<opacity>)", "Expands to a color from the tailwind CSS color palette. Opacity can be any value between 0-100. Ex: $color(red-500/80)."),
		MacroDescription("$<spacing>", "Expands to `calc(var(--spacing) * <spacing>`. Ex: `padding: $5;` or `padding: $3.5;`"),
		MacroDescription("$dark", "Expands to: `(prefers-color-scheme: dark)`"),
		MacroDescription("$light", "Expands to: `(prefers-color-scheme: light)`"),
		MacroDescription("$xs-", "Expands to: `screen and (max-width: 639px)`"),
		MacroDescription("$sm-", "Expands to: `screen and (max-width: 767px)`"),
		MacroDescription("$md-", "Expands to: `screen and (max-width: 1023px)`"),
		MacroDescription("$lg-", "Expands to: `screen and (max-width: 1279px)`"),
		MacroDescription("$xl-", "Expands to: `screen and (max-width: 1535px)`"),
		MacroDescription("$sm", "Expands to: `screen and (min-width: 640px)`"),
		MacroDescription("$md", "Expands to: `screen and (min-width: 768px)`"),
		MacroDescription("$lg", "Expands to: `screen and (min-width: 1024px)`"),
		MacroDescription("$xl", "Expands to: `screen and (min-width: 1280px)`"),
		MacroDescription("$xx", "Expands to: `screen and (min-width: 1536px)`"),

		InlineStyleComponent(),
	).Render(w)
}

func InlineStyleComponent() Node {
	return P(
		InlineStyle("$me { margin-top: $5; color: $color(blue-500); } @media $md- { $me { color: $color(red-500); padding: $5; } }"),
		InlineStyle("$me { font-size: var(--text-lg); }"),
		Text("You can call the `InlineStyle` macro as many times as you want on the same element. "),
		Text("Each macro call will get its own unique HTML attribute, unless it is a duplicate."),
	)
}

func MacroDescription(macro string, description string) Node {
	return P(
		InlineStyle("$me { margin-bottom: $3; display: flex; flex-direction: column;}"),
		Pre(InlineStyle("$me { color: $color(pink-600);}"), Code(Text(macro))),
		Text(" " + description),
	)
}