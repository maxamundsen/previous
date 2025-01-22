package app

import (
	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/auth"
	"previous/middleware"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
func DashboardController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	DashboardView(*identity).Render(w)
}

func DashboardView(identity auth.Identity) Node {
	return AppLayout("Dashboard", identity,
		H5(Class("font-bold"), Text("Welcome back, "), Text(identity.User.Firstname+" "+identity.User.Lastname), Text(".")),
		P(
			Text("This page requires a login!"),
		),
		Br(),
		P(
			Text("Here are some icons:"),
		),
		Span(Class("text-neutral-500"),
			Icon(ICON_GO, 24),
			Icon(ICON_GLOBE, 24),
			Icon(ICON_HTMX, 24),
			Icon(ICON_GITHUB, 24),
			Icon(ICON_X_DOT_COM, 24),
		),
		If(identity.User.PermissionAdmin != 0,
			Div(Class("mt-10 p-10 bg-white border border-neutral-200 shadow"),
				P(Class("font-bold text-red-600"), Text("Admin only")),
				P(Text("You can only see this if you have the admin permission")),
			),
		),
	)
}
