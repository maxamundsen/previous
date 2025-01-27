package app

import (
	"fmt"
	"previous/.metagen/pageinfo"
	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/middleware"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
func DashboardPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	// You can get page info from the PageInfoMap, a mapping between urls and PageInfo.
	// Since this handler handles requests to that URL, we can guarantee that the incoming URL matches the correct PageInfo
	thisPageInfo := pageinfo.PageInfoMap[r.URL.Path]

	func() Node {
		return AppLayout("Dashboard", *identity,
			H5(Class("font-bold"), Text("Welcome back, "), Text(identity.User.Firstname+" "+identity.User.Lastname), Text(".")),
			P(
				Text("This page requires a login!"),
			),
			Br(),
			P(
				Text("Here are some icons:"),
			),
			Span(Class("text-cyan-400"),
				Icon(ICON_GO, 24),
			),
			Span(Class("ml-3 text-orange-500"),
				Icon(ICON_RSS, 24),
			),
			Span(Class("ml-3 text-blue-500"),
				Icon(ICON_HTMX, 24),
			),
			Span(Class("ml-3 text-neutral-900"),
				Icon(ICON_GITHUB, 24),
			),
			Span(Class("ml-3 text-black"),
				Icon(ICON_X_DOT_COM, 24),
			),
			Br(),
			Br(),
			P(
				Text("Page handlers support \"reflection.\" For example, the page you are reading right now is defined in:"),
				Br(),
				B(Text(thisPageInfo.FileDef)),
				Br(),
				Text("the URL is:"),
				Br(),
				B(Text(thisPageInfo.Url)),
				Br(),
				Text("and has the following middleware:"),
				Br(),
				B(Text(fmt.Sprintf("%+v", thisPageInfo.Middleware))),
			),
			If(identity.User.PermissionAdmin != 0,
				Div(Class("mt-10 p-10 bg-white border border-neutral-200 shadow"),
					P(Class("font-bold text-red-600"), Text("Admin only")),
					P(Text("You can only see this if you have the admin permission")),
				),
			),
		)
	}().Render(w)
}
