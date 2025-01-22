package docs

import (
	"os"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"

	. "previous/components"

	"github.com/gomarkdown/markdown"

	"net/http"
	"path"
)

type Document struct {
	DisplayId    int
	SubDisplayId int
	Title        string
	Slug         string
	SubList      []Document
}

var DocList []Document

func RegisterDocPage(doc Document) {
	doc.DisplayId = len(DocList) + 1

	for i := range doc.SubList {
		doc.SubList[i].DisplayId = len(DocList) + 1
	}

	DocList = append(DocList, doc)
}

func RegisterDocumentation() {
	RegisterDocPage(Document{
		Title: "Applications",
		Slug:  "applications",
	})

	RegisterDocPage(Document{
		Title: "Metaprogram / Build System",
		Slug:  "metaprogram",
	})

	RegisterDocPage(Document{
		Title: "Configuration",
		Slug:  "configuration",
	})

	RegisterDocPage(Document{
		Title: "Routing",
		Slug:  "routing",
	})

	RegisterDocPage(Document{
		Title: "Controllers",
		Slug:  "controllers",
	})

	RegisterDocPage(Document{
		Title: "Views",
		SubList: []Document{
			{Title: "Component System", Slug: "component-system"},
			{Title: "Interactivity", Slug: "interactivity"},
			{Title: "Markdown Content", Slug: "markdown-content"},
			{Title: "Tailwind Integration", Slug: "tailwind-integration"},
			{Title: "Icons", Slug: "icons"},
		},
	})

	RegisterDocPage(Document{
		Title: "Auth Flow",
		SubList: []Document{
			{Title: "Authentication", Slug: "authentication"},
			{Title: "Authorization", Slug: "authorization"},
		},
	})

	RegisterDocPage(Document{
		Title: "Middleware",
		SubList: []Document{
			{Title: "Middleware Chains", Slug: "middleware-chains"},
			{Title: "Auth / Identity", Slug: "middleware-identity"},
			{Title: "Sessions", Slug: "middleware-sessions"},
		},
	})

	RegisterDocPage(Document{
		Title: "Extras / Helpers",
		SubList: []Document{
			{Title: "SMTP Client", Slug: "smtp-client"},
			{Title: "View Helpers / Formatters", Slug: "view-helpers"},
			{Title: "Financial Helpers", Slug: "financial-helpers"},
		},
	})

	RegisterDocPage(Document{
		Title: "Examples",
		Slug:  "examples",
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
	path := "./README.md"
	mdContent, _ := os.ReadFile(path)

	html := markdown.ToHTML(mdContent, nil, nil)

	DocView("Overview", 0, string(html)).Render(w)
}

func DocController(w http.ResponseWriter, r *http.Request) {
	doc := FindDocumentationByURL(path.Base(r.URL.Path))
	path := "./docs/" + doc.Slug + ".md"
	mdContent, _ := os.ReadFile(path)
	html := markdown.ToHTML(mdContent, nil, nil)

	DocView(doc.Title, doc.DisplayId, string(html)).Render(w)
}

func DocView(title string, displayId int, html string) Node {
	return DocLayout(title, displayId,
		Raw(html),
	)
}

func DocLayout(title string, displayId int, children ...Node) Node {
	return RootLayout(title+" | Previous Documentation",
		Body(Attr("x-data", "{ mobileMenu: false }"), Attr("hx-swap", "innerHTML show:unset"),
			Button(Attr("x-on:click", "mobileMenu = !mobileMenu"), Type("button"), Class("inline-flex items-center p-2 mt-2 ms-3 text-sm text-neutral-800 sm:hidden hover:bg-neutral-100 focus:outline-none focus:ring-2 focus:ring-neutral-200"),
				Span(Class("sr-only"), Text("Open sidebar")),
				Icon(ICON_MENU, 24),
			),
			Aside(Class("border-1 shadow-sm bg-zinc-100 fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0 overflow-y-auto"),
				Div(Class("px-4 overflow-y-auto py-6"),
					A(Href("/"), Img(Class("mx-auto h-16 w-auto"), Src("/images/logo.svg"), Alt("Previous"))),
					H5(Class("mt-3 mb-5 text-center text-neutral-800 "), Text("Previous: Codebase Documentation")),
					Ul(Class("mt-6 space-y-1"),
						A(Href("/docs"), Classes{"block px-4 py-2 text-sm font-medium text-neutral-800 hover:bg-neutral-200": true, "bg-neutral-200": displayId == 0}, Text("Overview")),
						Map(DocList, func(doc Document) Node {
							if len(doc.SubList) > 0 {
								return Li(
									Details(Class("group [&_summary::-webkit-details-marker]:hidden"), If(doc.DisplayId == displayId, Attr("open")),
										Summary(Class("flex cursor-pointer items-center justify-between px-4 py-2 text-neutral-800 hover:bg-neutral-200"),
											Span(Class("text-sm font-medium"), Text(doc.Title)),
											Span(Class("shrink-0 transition duration-300 group-open:-rotate-180"),
												Icon(ICON_CHEVRON_UP, 16),
											),
										),
										Ul(Class("mt-2 space-y-1 px-4"),
											Map(doc.SubList, func(subdoc Document) Node {
												return Li(
													A(Href("/docs/"+subdoc.Slug), Classes{"block px-4 py-2 text-sm font-medium text-neutral-800": true, "hover:bg-neutral-200": title != subdoc.Title, "bg-neutral-200": title == subdoc.Title}, Text(subdoc.Title)),
												)
											}),
										),
									),
								)
							} else {
								return Li(
									A(Href("/docs/"+doc.Slug), Classes{"block px-4 py-2 text-sm font-medium text-neutral-800 hover:bg-neutral-200": true, "bg-neutral-200": displayId == doc.DisplayId}, Text(doc.Title)),
								)
							}
						}),
					),
				),
			),
			Div(Class("m-5 p-10 sm:ml-72 prose prose-pre:rounded-none prose-pre:text-neutral-700 prose-pre:bg-neutral-100 max-w-none bg-white rose-a:text-neutral-800  prose-headings:text-neutral-950"),
				Group(children),
			),
		),
		Script(Raw(`
			hljs.highlightAll();
		`)),
	)
}
