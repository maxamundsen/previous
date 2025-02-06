package examples

import (
	. "maragu.dev/gomponents"
	// . "maragu.dev/gomponents/html"
	"net/http"
	. "previous/components"
	"previous/middleware"
	. "previous/pages/app"
)

// @Identity
// @Protected
// @CookieSession
func UIElementsPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	const (
		MODAL1 = "modal1"
		MODAL2 = "modal2"
	)

	// Alpine stores are used to handle global state.
	// In this case, we initialize a store containing two keys: "modal1" and "modal2", each mapping to the value "false".
	// When AlpineStoreInit is called, it generates the javascript that binds

	store := make(AlpineStore)

	store[MODAL1] = "false"
	store[MODAL2] = "false"

	func() Node {
		return AppLayout("User Interface Elements", *identity,
			AlpineStoreInit(store),

			ModalActuator(MODAL1, ButtonBlue(Text("Open Modal 1"))),
			ModalActuator(MODAL2, ButtonRed(Text("Open Modal 2"))),

			Modal(
				MODAL1,
				"Test Modal",
				Text("Hello!"),
				[]Node{
					ButtonGray(Text("Close")),
					ButtonRed(Text("Click this!")),
				},
			),

			Modal(
				MODAL2,
				"Another Modal",
				Text(LOREM_IPSUM),
				[]Node{
					ButtonGray(Text("Close")),
				},
			),
		)
	}().Render(w)
}
