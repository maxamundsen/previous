package app

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"webdawgengine/middleware"
	"webdawgengine/models"

	"net/http"
)

func DashboardController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	DashboardView(*identity).Render(w)
}

func DashboardView(identity models.Identity) Node {
	return AppLayout("Dashboard", identity,
		H5(Class("font-bold"), Text("Welcome back, "), Text(identity.User.Firstname+" "+identity.User.Lastname), Text(".")),
		P(
			Text("This page requires a login!"),
		),
		If(identity.User.PermissionAdmin,
			Div(
				Div(Class("mt-10 p-10 bg-white border border-gray-200 rounded-lg shadow"),
					P(Class("font-bold text-red-600"), Text("Admin only")),
					P(Text("You can only see this if you have the admin permission")),
				),
			),
		),
	)
}
