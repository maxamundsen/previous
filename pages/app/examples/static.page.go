package examples

import (
	"net/http"
	"previous/.metagen/pageinfo"
	. "previous/components"
	"previous/middleware"
	. "previous/pages/app"
	"time"

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
					P(
						Text("Check the source code located at "),
						B(Text(this.FileDef())),
						Text(" and remove the line containing:"),
						Pre(Class("text-red-600"), Raw("// @Static")),
						Text(" to see the page in dynamic mode."),
					),
					Br(),
					P(
						InlineStyle("color: red;"),
						Text("You can run arbitary code at the page level, and the result will be baked in at compile time. For example, 1 + 1 = "),
						ToText(1+1),
						Text(". This page was generated at "), FormatDateTime(time.Now()), Text("."),
					),
					Div(Class("mt-10 p-10 bg-white border border-neutral-200 shadow"),
						P(Class("font-bold text-red-600"), Text("IMPORTANT")),
						P(
							Text("It is important to understand that code generated at compile time runs on the build machine, against the current environment variables. "),
							Text("This may be obvious, but it is important to always keep this in mind when building static pages. "),
							Text("Static pages can still execute code, and access global resources, but ONLY at the time of compilation. "),
						),
						Br(),
						P(
							Text("Despite the power of compile-time execution, you should probably avoid "),
						),
					),
				},
				P(Text("Page is running in dynamic mode. Add @Static directive to page handler function to see the rest of the contents.")),
			),
		)
	}().Render(w)
}
