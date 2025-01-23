package examples

import (
	. "maragu.dev/gomponents"

	. "previous/components"
	. "previous/pages/app"

	// "previous/repository"

	"previous/auth"
	"previous/middleware"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
func PaginationPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	PaginationView(*identity).Render(w)
}

func PaginationView(identity auth.Identity) Node {
	return AppLayout("Pagination", identity,
		HxLoad("/app/examples/orders-hx"),
	)
}