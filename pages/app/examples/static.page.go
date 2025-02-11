package examples

import (
	"net/http"
	"previous/.metagen/pageinfo"
	"previous/middleware"
	. "previous/pages/app"
	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// @Identity
// @Protected
// @CookieSession
// @Static
func StaticPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	this := pageinfo.Reflect(r)

	func() Node {
		return AppLayout("Static Page", *identity,
			IfElse(this.IsStatic(),
				Group{
					P(Text("This page is static and has been pre-compiled to HTML.")),
					Br(),
					P(Text("Check the source code located at "), B(Text(this.FileDef())), Text(" and uncomment the line containing:"),
						Pre(Class("text-red-600"), Raw("@Static")),
					),
				},
				P(Text("Page is running in dynamic mode. Add @Static directive to page handler function.")),
			),
		)
	}().Render(w)
}
