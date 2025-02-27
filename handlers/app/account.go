package app

import (
	"previous/middleware"

	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"net/http"
)

type accountTableItem struct {
	Property string
	Value interface{}
}

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	cols := []string {
		"Property",
		"Value",
	}

	entries := []accountTableItem {
		{ "UserId", identity.User.ID },
		{ "Username", identity.User.Username },
		{ "Name", identity.User.Username + " " +  identity.User.Lastname },
		{ "Email", identity.User.Email },
		{ "Last Login", identity.User.LastLogin },
	}

	func() Node {
		return AppLayout("Account", *identity,
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
