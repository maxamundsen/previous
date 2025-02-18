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
	this := pageinfo.Reflect(r)

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
			Span(InlineStyle("$this{color: var(--color-cyan-400);}"),
				Icon(ICON_GO, 24),
			),
			Span(InlineStyle("@media $lg-{@media $dark{$this{color: red;}}} $this{margin-left: $(3); color: var(--color-orange-400);}"),
				Icon(ICON_RSS, 24),
			),
			Span(InlineStyle("$this { margin-left: $(3); color: var(--color-blue-400); }"),
				Icon(ICON_HTMX, 24),
			),
			Span(InlineStyle("$this {padding-left: $(3); padding-right: $(3); color: var(--color-neutral-900);}"),
				Icon(ICON_GITHUB, 24),
			),
			Span(InlineStyle("$this{ margin-left: $(3); color: var(--color-black); }"),
				Icon(ICON_X_DOT_COM, 24),
			),
			Br(),
			Br(),
			P(
				Text("Page handlers support \"reflection.\" For example, the page you are reading right now is defined in:"),
				Br(),
				B(Text(this.FileDef())),
				Br(),
				Text("the URL is:"),
				Br(),
				B(Text(this.Url())),
				Br(),
				Text("and has the following middleware:"),
				Br(),
				B(Text(fmt.Sprintf("%+v", this.Middleware()))),
			),
			If(identity.User.PermissionAdmin != 0,
				Div(InlineStyle("$this { margin-top: $(10); padding: $(10); background-color: var(--color-white); border: 1px solid var(--color-neutral-200); box-shadow: var(--shadow-md); }"),
					P(Class("font-bold text-red-600"), Text("Admin only")),
					P(Text("You can only see this if you have the admin permission")),
				),
			),
		)
	}().Render(w)
}
