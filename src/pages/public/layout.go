package public

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"
)

func PublicLayout(title string, children ...Node) Node {
	return Root(title,
		Div(Class("container"),
			Group(children),
		),
	)
}
