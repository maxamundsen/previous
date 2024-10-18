package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func Root(title string, children ...Node) Node {
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("/lib/bootstrap/bootstrap.min.css")),
			Link(Rel("stylesheet"), Href("/lib/bootstrap-icons/bootstrap-icons.min.css")),
			Link(Rel("stylesheet"), Href("/css/style.css")),
			Script(Src("/lib/htmx/htmx.min.js")),
			Script(Src("/lib/bootstrap/bootstrap.bundle.min.js")),
			Script(Src("/js/index.js")),
		},
		Body: Group(children),
	})
}

func EmailRoot(children ...Node) Node {
	return HTML5(HTML5Props{
		Language: "en",
		Body:     Group(children),
	})
}
