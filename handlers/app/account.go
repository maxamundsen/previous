package app

import (
	"previous/middleware"

	. "previous/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"net/http"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)


	type accountTableItem struct {
		Property string
		Value interface{}
	}

	cols := []string {
		"Property",
		"Value",
	}

	entries := []accountTableItem {
		{ "UserId", identity.User.ID },
		{ "Username", identity.User.Username },
		{ "Name", identity.User.Firstname + " " +  identity.User.Lastname },
		{ "Email", identity.User.Email },
		{ "Last Login", identity.User.LastLogin },
	}

	func() Node {
		return AppLayout("Account", LAYOUT_SECTION_ACCOUNT, *identity, session,
			AutoTableLite(
				cols,
				entries,
				func(item accountTableItem) Node {
					return Tr(
						Td(B(Text(item.Property))),
						Td(ToText(item.Value)),
					)
				},
				AutoTableOptions{
					BorderX: true,
				},
			),
		)
	}().Render(w)
}
