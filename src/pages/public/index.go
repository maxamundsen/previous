package public

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
	. "webdawgengine/pages/components"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	IndexView().Render(w)
}

func IndexView() Node {
	return PublicLayout("WebDawgEngine",
		Div(Class("flex pt-10 items-center"),
			A(Href("/"), Img(Class("hidden sm:block me-24"), Src("/images/logo.svg"), Width("150"), Height("150"))),
			Div(
				H1(Class("mt-7 text-balance text-7xl font-semibold text-gray-900"), Text("WebDawgEngine")),
				H4(Class("pt-5 text-balance text-3xl font-semibold text-gray-700"), Text("A system for building modern websites and web applications without unnecessary abstractions and complexities.")),
			),
		),
		A(Class("mt-7 group relative inline-flex items-center overflow-hidden rounded-lg bg-red-600 px-8 py-3 text-white focus:outline-none focus:ring hover:bg-red-800 active:bg-red-800"), Href("/docs"),
			Span(Class("absolute -start-full transition-all group-hover:start-4"),
				Icon("arrow-right", 24),
			),
			Span(Class("text-sm font-medium transition-all group-hover:ms-4"), Text("View the Documentation!")),
		),
	)
}
