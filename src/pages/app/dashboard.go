package app

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"webdawgengine/database"
	"webdawgengine/middleware"
	"webdawgengine/models"

	"net/http"
)

func DashboardController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	user, _ := database.FetchUserById(identity.UserId)

	DashboardView(user).Render(w)
}

func DashboardView(user models.User) Node {
	return AppLayout("Dashboard",
		H5(Text("Welcome back, "), Text(user.Firstname+" "+user.Lastname), Text(".")),
	)
}
