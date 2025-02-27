package app

import (
	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/middleware"

	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	func() Node {
		return AppLayout("Dashboard", *identity,
			H5(Class("font-bold"), Text("Welcome back, "), Text(identity.User.Firstname+" "+identity.User.Lastname), Text(".")),
			P(
				Text("This page requires a login!"),
				Text("previous/pages"),
			),
			Br(),
			P(
				Text("Here are some icons:"),
			),
			Span(InlineStyle("$me{color: $color(cyan-400);}"),
				Icon(ICON_GO, 24),
			),
			Span(InlineStyle("@media $lg-{ @media $dark{ $me{color: red;}}} $me{margin-left: $3; color: $color(orange-400);}"),
				Icon(ICON_RSS, 24),
			),
			Span(InlineStyle("$me{ margin-left: $3; color: $color(blue-400); }"),
				Icon(ICON_HTMX, 24),
			),
			Span(InlineStyle("$me{padding-left: $3; padding-right: $3; color: $color(neutral-900);}"),
				Icon(ICON_GITHUB, 24),
			),
			Span(InlineStyle("$me{ margin-left: $3; color: $color(black); }"),
				Icon(ICON_X_DOT_COM, 24),
			),
			If(identity.User.PermissionAdmin != 0,
				Div(InlineStyle("$me{ margin-top: $10; padding: $10; background-color: $color(white); border: 1px solid $color(neutral-200); box-shadow: var(--shadow-md); }"),
					P(InlineStyle("$me { font-weight: var(--font-weight-bold); color: $color(red-600); }"), Text("Admin only")),
					P(Text("You can only see this if you have the admin permission")),
				),
			),
		)
	}().Render(w)
}
