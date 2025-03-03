package app

import (
	. "previous/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/middleware"

	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	layout := r.URL.Query().Get("layout")

	if layout == "vertical" {
		session["APP_LAYOUT_VERTICAL"] = true
		middleware.PutSessionCookie(w, r, session)
	} else if layout == "horizontal" {
		session["APP_LAYOUT_VERTICAL"] = false
		middleware.PutSessionCookie(w, r, session)
	}

	func() Node {
		iconContainer := func (c ...Node) Node {
			return Span(InlineStyle("$me{padding: $2; background: $color(white); border: 1px solid $color(neutral-200); border-radius: var(--radius-sm); color: $color(orange-400);}"),
				Group(c),
			)
		}

		return AppLayout("Dashboard", LAYOUT_SECTION_DASHBOARD, *identity, session,
			H5(
				InlineStyle("$me { font-size: var(--text-lg); }"),
				Text("Welcome back, "), B(Text(identity.User.Firstname+" "+identity.User.Lastname)),
				Text("."),
			),

			Br(),

			P(
				Text("This page requires a login!"),
			),

			Br(),

			P(
				Form(
					Action(""),
					Method("GET"),
					Input(Type("hidden"), Name("layout"),
						IfElse(session["APP_LAYOUT_VERTICAL"] == true,
							Value("horizontal"),
							Value("vertical"),
						),
					),

					ButtonUI(Type("submit"), Text("Click to toggle layout")),
				),
			),

			Br(),

			P(
				Text("Here are some icons:"),
			),

			Br(),

			Div(InlineStyle("$me { display: flex; gap: $4; align-items: center; margin-bottom: $10; }"),
				iconContainer(InlineStyle("$me{ color: $color(cyan-400);}"),
					Icon(ICON_GO, 24),
				),
				iconContainer(InlineStyle("$me{ color: $color(orange-500);}"),
					Icon(ICON_RSS, 24),
				),
				iconContainer(InlineStyle("$me{ color: $color(blue-600); }"),
					Icon(ICON_HTMX, 24),
				),
				iconContainer(InlineStyle("$me{ color: $color(neutral-900);}"),
					Icon(ICON_GITHUB, 24),
				),
				iconContainer(InlineStyle("$me{ color: $color(black); }"),
					Icon(ICON_X_DOT_COM, 24),
				),
				iconContainer(InlineStyle("$me{ color: $color(black); }"),
					Icon(ICON_XAI_GROK, 24),
				),
				iconContainer(InlineStyle("$me{ color: #006600 }"),
					Icon(ICON_4CH, 24),
				),
			),

			If(identity.User.PermissionAdmin != 0,
				Card(
					P(InlineStyle("$me { font-weight: var(--font-weight-bold); color: $color(red-600); }"), Text("Admin only")),
					P(Text("You can only see this if you have the admin permission")),
				),
			),
		)
	}().Render(w)
}
