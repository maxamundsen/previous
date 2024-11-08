package docs

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "maragu.dev/gomponents/components"


	. "webdawgengine/pages/components"
	"github.com/gomarkdown/markdown"

	"net/http"
	"path"
	"io/ioutil"
)

type Document struct {
	DisplayId int
	SubDisplayId int
	Title string
	Slug string
	SubList []Document
}

var DocList []Document

func RegisterDocPage(doc Document) {
	doc.DisplayId = len(DocList) + 1

	for i, _ := range doc.SubList {
		doc.SubList[i].DisplayId = len(DocList) + 1
	}

	DocList = append(DocList, doc)
}

func RegisterDocumentation() {
	RegisterDocPage(Document{
		Title: "Pre-requisites",
		Slug: "prerequisites",
	})

	RegisterDocPage(Document{
		Title: "Getting Started",
		Slug: "getting-started",
	})

	RegisterDocPage(Document{
		Title: "Build System",
		Slug: "build-system",
	})

	RegisterDocPage(Document{
		Title: "Program Entrypoint",
		Slug: "program-entrypoint",
	})

	RegisterDocPage(Document{
		Title: "Configuration",
		Slug: "configuration",
	})

	RegisterDocPage(Document{
		Title: "Routing",
		SubList: []Document{
			{ Title: "Static Routing", Slug: "static-routing" },
			{ Title: "Dynamic Routing", Slug: "dynamic-routing" },
		},
	})

	RegisterDocPage(Document{
		Title: "Controllers",
		SubList: []Document {
			{ Title: "Page Controllers", Slug: "page-controllers" },
			{ Title: "API Controllers", Slug: "api-controllers" },
		},
	})

	RegisterDocPage(Document{
		Title: "Views",
		SubList: []Document {
			{ Title: "Generating HTML", Slug: "generating-html" },
			{ Title: "Component System", Slug: "component-system" },
			{ Title: "Organizing Components", Slug: "organizing-components" },
			{ Title: "Markdown Content", Slug:"markdown-content" },
		},
	})

	RegisterDocPage(Document{
		Title: "Styling",
		SubList: []Document{
			{ Title: "Tailwind Integration", Slug: "tailwind-integration" },
			{ Title: "Icons", Slug: "icons" },
		},
	})

	RegisterDocPage(Document{
		Title: "Interactivity",
		SubList: []Document{
			{ Title: "HTMX", Slug: "htmx" },
			{ Title: "Alpine.js", Slug: "alpine" },
			{ Title: "Component Integration", Slug: "interactive-components" },
		},
	})

	RegisterDocPage(Document{
		Title: "Database Interaction",
		SubList: []Document{
			{ Title: "DB Connections", Slug: "db-connections" },
			{ Title: "SQL Compiler", Slug: "sql-compiler" },
		},
	})

	RegisterDocPage(Document{
		Title: "Auth Flow",
		SubList: []Document{
			{  Title: "Authentication", Slug: "authentication" },
			{  Title: "Authorization", Slug: "authorization" },
		},
	})

	RegisterDocPage(Document{
		Title: "Middleware",
		SubList: []Document{
			{  Title: "Middleware Chains", Slug: "middleware-chains" },
			{  Title: "Auth / Identity", Slug: "middleware-identity" },
			{  Title: "Sessions", Slug: "middleware-sessions" },
		},
	})

	RegisterDocPage(Document{
		Title: "Editor Support",
		Slug: "editor-support",
	})

	RegisterDocPage(Document{
		Title: "Examples",
		Slug: "examples",
	})
}

func FindDocumentationByURL(url string) Document {
	doc := Document{}

	for _, v := range DocList {
		if len(v.SubList) > 0 {
			for _, k := range v.SubList {
				if k.Slug == url {
					doc = k
				}
			}
		} else {
			if v.Slug == url {
				doc = v
			}
		}
	}

	return doc
}

func IndexController(w http.ResponseWriter, r *http.Request) {
	path := "./pages/docs/overview.md"
	mdContent, _ := ioutil.ReadFile(path)
	html := markdown.ToHTML(mdContent, nil, nil)

	DocView("Overview", 0, string(html)).Render(w)
}

func DocController(w http.ResponseWriter, r *http.Request) {
	doc := FindDocumentationByURL(path.Base(r.URL.Path))
	path := "./pages/docs/" + doc.Slug + ".md"
	mdContent, _ := ioutil.ReadFile(path)
	html := markdown.ToHTML(mdContent, nil, nil)

	DocView(doc.Title, doc.DisplayId, string(html)).Render(w)
}

func DocView(title string, displayId int, html string) Node {
	return DocLayout(title, displayId,
		Raw(html),
	)
}

func DocLayout(title string, displayId int, children ...Node) Node {
	return Root(title + " | WebDawgEngine Documentation",
		Body(Attr("x-data", "{ mobileMenu: false }"), Attr("hx-boost", "true"), Attr("hx-swap", "innerHTML show:unset"), Class("bg-gray-50"),
			Button(Attr("x-on:click", "mobileMenu = !mobileMenu"), Type("button"), Class("inline-flex items-center p-2 mt-2 ms-3 text-sm text-gray-500 rounded-lg sm:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"),
				Span(Class("sr-only"), Text("Open sidebar")),
				Icon("menu", 24),
			),
			Aside(Class("border-r border-gray-200 shadow-sm bg-white fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0 overflow-y-auto"),
				Div(Class("px-4 overflow-y-auto py-6"),
					A(Attr("hx-boost", "false"), Href("/"), Img(Class("mx-auto h-10 w-auto"), Src("/images/logo.svg"), Alt("WebDawgEngine"))),
					H5(Class("mt-3 mb-5 text-center text-gray-900 dark:text-white"), Text("WebDawgEngine Documentation")),
					Ul(Class("mt-6 space-y-1"),
						IfElse(displayId == 0,
							A(Href("/docs"), Class("block rounded-lg px-4 py-2 text-sm font-medium text-gray-800 bg-gray-100"), Text("Overview")),
							A(Href("/docs"), Class("block rounded-lg px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-100"), Text("Overview")),
						),
						Map(DocList, func(doc Document) Node {
							if len(doc.SubList) > 0 {
								return Li(
									Details(Class("group [&_summary::-webkit-details-marker]:hidden"), If(doc.DisplayId == displayId, Attr("open")),
										Summary(Class("flex cursor-pointer items-center justify-between rounded-lg px-4 py-2 text-gray-500 hover:bg-gray-100 hover:text-gray-700"),
											Span(Class("text-sm font-medium"), Text(doc.Title)),
											Span(Class("shrink-0 transition duration-300 group-open:-rotate-180"),
												Icon("chevron-down", 16),
											),
										),
										Ul(Class("mt-2 space-y-1 px-4"),
											Map(doc.SubList, func(subdoc Document) Node {
												return Li(
													A(Href("/docs/" + subdoc.Slug), Classes{"block rounded-lg px-4 py-2 text-sm font-medium": true, "text-gray-500 hover:bg-gray-100": title != subdoc.Title, "bg-gray-100 text-gray-800": title == subdoc.Title}, Text(subdoc.Title)),
												)
											}),
										),
									),
								)
							} else {
								return Li(
									IfElse(doc.DisplayId == displayId,
										A(Href("/docs/" + doc.Slug), Class("block bg-gray-100 rounded-lg px-4 py-2 text-sm font-medium text-gray-800"), Text(doc.Title)),
										A(Href("/docs/" + doc.Slug), Class("block rounded-lg px-4 py-2 text-sm font-medium text-gray-500 hover:bg-gray-100"), Text(doc.Title)),
									),
								)
							}
						}),
					),
				),
			),
			Div(Attr("hx-boost", "false"), Class("m-5 rounded-xl p-10 sm:ml-72 prose bg-white ring-1 ring-inset ring-gray-200 prose-img:rounded-xl prose-a:text-red-800"),
				Group(children),
			),
		),
		Script(Raw(`
			hljs.highlightAll();
		`)),
	)
}