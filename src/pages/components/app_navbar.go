package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AppNavbar() Node {
	return Nav(Class("navbar navbar-expand-lg bg-light mb-3"), Data("bs-theme", "light"),
		Div(Class("container-fluid"),
			A(Class("navbar-brand"), Href("/"),
				Img(Src("/images/logo.png"), Alt("Logo"), Height("30"), Width("auto")),
			),
			Button(Class("navbar-toggler"), Type("button"), Data("bs-toggle", "collapse"), Data("bs-target", "#navbarColor03"), Aria("controls", "navbarColor03"), Aria("expanded", "false"), Aria("label", "Toggle navigation"),
				Span(Class("navbar-toggler-icon")),
			),
			Div(Class("collapse navbar-collapse"), ID("navbarColor03"),
				Ul(Class("navbar-nav me-auto"),
					Li(Class("nav-item"),
						A(Class("nav-link"), Href("/app/dashboard"), Text("Dashboard")),
					),
					Li(Class("nav-item dropdown"),
						A(Class("nav-link dropdown-toggle"), Data("bs-toggle", "dropdown"), Href("#"), Role("button"), Aria("haspopup", "true"), Aria("expanded", "false"), Text("Examples")),
						Div(Class("dropdown-menu"),
							A(Class("dropdown-item"), Href("/app/examples/api-fetch"), Text("API Fetch")),
							A(Class("dropdown-item"), Href("/app/examples/htmx"), Text("HTMX")),
							A(Class("dropdown-item"), Href("/app/examples/upload"), Text("File Upload")),
							A(Class("dropdown-item"), Href("/app/examples/smtp"), Text("SMTP client")),
							A(Class("dropdown-item"), Href("/404"), Text("Error 404 Page")),
						),
					),
					Li(Class("nav-item"),
						A(Class("nav-link"), Href("/app/api-tester"), Text("API Tester")),
					),
					Li(Class("nav-item"),
						A(Class("nav-link"), Href("/app/account"), Text("Account")),
					),
				),
				Div(Class("d-flex"),
					A(Href("/auth/logout"),
						Button(Class("btn btn-outline-secondary my-2 my-sm-0"), I(Class("bi bi-box-arrow-right"))),
					),
				),
			),
		),
	)
}
