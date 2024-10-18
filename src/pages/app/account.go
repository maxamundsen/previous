package app

import (
	"webdawgengine/database"
	"webdawgengine/middleware"
	"webdawgengine/models"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"net/http"
)

func AccountController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	user, _ := database.FetchUserById(identity.UserId)

	AccountView(user).Render(w)
}

func AccountView(user models.User) Node {
	return AppLayout("Account",
		Div(Class("table-responsive"), Style("table-layout: fixed"),
			Table(Class("table"),
				THead(),
				TBody(
					Tr(
						Th(Text("UserID")),
						Td(ToText(user.Id)),
					),
					Tr(
						Th(Text("Username")),
						Td(ToText(user.Username)),
					),
					Tr(
						Th(Text("Name")),
						Td(ToText(user.Firstname+" "+user.Lastname)),
					),
					Tr(
						Th(Text("Email")),
						Td(ToText(user.Email)),
					),
				),
			),
		),
	)
}
