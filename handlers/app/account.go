package app

import (
	"previous/middleware"

	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"net/http"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	func() Node {
		return AppLayout("Account", *identity,
			Div(Class("table-responsive"),
				H4(Class("mb-3 text-lg font-bold"), Text("General Account Information")),
				Hr(),
				TableTW(
					THead(),
					TBody(
						Tr(
							ThTW(Text("UserID")),
							TdTW(ToText(identity.User.ID)),
						),
						Tr(
							ThTW(Text("Username")),
							TdTW(ToText(identity.User.Username)),
						),
						Tr(
							ThTW(Text("Name")),
							TdTW(ToText(identity.User.Firstname+" "+identity.User.Lastname)),
						),
						Tr(
							ThTW(Text("Email")),
							TdTW(ToText(identity.User.Email)),
						),
						Tr(
							ThTW(Text("Last Login")),
							TdTW(Text(identity.User.LastLogin)),
						),
					),
				),
				Br(),
				H4(Class("mb-3 text-lg font-bold"), Text("Permissions")),
				Hr(),
				TableTW(
					THead(
						Tr(
							ThTW(Text("Permission")),
							ThTW(Text("Value")),
						),
					),
					TBodyTW(
						Tr(
							TdTW(Text("Admin")),
							TdTW(Input(Type("checkbox"), If(identity.User.PermissionAdmin != 0, Checked()), Disabled())),
						),
					),
				),
			),
		)
	}().Render(w)
}
