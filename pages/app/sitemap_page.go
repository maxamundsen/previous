package app

import (
	"fmt"
	"net/http"
	"previous/.metagen/pageinfo"
	. "previous/components"
	"previous/middleware"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// @Identity
// @Protected
// @CookieSession
func SitemapPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	func() Node {
		return AppLayout("Sitemap", *identity,
			Map(pageinfo.PageInfoList, func(p pageinfo.PageInfo) Node {
				return Group{
					PageLink(p.Url(), Text(p.Url()), false),
					Div(Class("ml-3"),
						Text("Source: "), Text(p.FileDef()),
						Br(),
						Text("Active Middleware: "), Text(fmt.Sprintf("%+v", p.Middleware())),
					),
					Br(),
				}
			}),
		)
	}().Render(w)
}
