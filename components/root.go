package components

import (
	"previous/config"
	"previous/crypt"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func RootLayout(title string, children ...Node) Node {
	// automatically invalidates cached css when file hash changes
	css_hash, err := crypt.QuickFileHash("./wwwroot/css/style.css")
	if err != nil {
		return Text("Error hashing style.css")
	}

	// automatically invalidates cached js when file hash changes
	js_hash, err := crypt.QuickFileHash("./wwwroot/js/index.js")
	if err != nil {
		return Text("Error hashing index.js")
	}

	return Doctype(
		HTML(Class("h-full"),
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleEl(Text(title)),
				Meta(Name("description"), Content("Previous")),

				Link(Rel("apple-touch-icon"), Attr("sizes", "180x180"), Href("/apple-touch-icon.png")),
				Link(Rel("icon"), Type("image/png"), Attr("sizes", "32x32"), Href("/favicon-32x32.png")),
				Link(Rel("icon"), Type("image/png"), Attr("sizes", "16x16"), Href("/favicon-16x16.png")),
				Link(Rel("manifest"), Href("/site.webmanifest")),

				Link(Rel("stylesheet"), Href("/fonts/inter.css")),
				Link(Rel("stylesheet"), Href("/fonts/lora.css")),

				IfElse(TW_FALLBACK,
					Group{
						StyleEl(Type("text/tailwindcss")),
						Script(Src("/lib/tailwind/index.global.js")),
					},
					Link(Rel("stylesheet"), Href("/css/style.css?v="+css_hash)),
				),

				Link(Rel("stylesheet"), Href("/lib/highlight/default.min.css")),

				// use minified htmx only in prod
				IfElse(config.DEBUG,
					Script(Src("/lib/htmx/htmx.js")),
					Script(Src("/lib/htmx/htmx.min.js")),
				),

				Script(Src("/lib/alpine/alpine.min.js"), Defer()),
				Script(Src("/lib/highlight/highlight.min.js")),
				Script(Src("/js/index.js?v="+js_hash)),
			),
			Group(children), // expected to provide body
		),
	)
}

func EmailRoot(children ...Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				Meta(Name("description"), Content("Previous")),
			),
			Body(
				Group(children),
			),
		),
	)
}
