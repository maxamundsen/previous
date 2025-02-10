package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	"net/http"
	"previous/.metagen/pageinfo"
	. "previous/components"
	"previous/middleware"
	. "previous/pages/app"
)

// @Identity
// @Protected
// @CookieSession
func UIPlaygroundPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	const (
		MODAL1 = "modal1"
		MODAL2 = "modal2"
	)

	// Sometimes you need to handle client-side global state.
	//     Ex: The page has a modal popup that can be triggered by many different actuators.

	// Alpine stores are used to handle javascript global state.
	// In this case, we initialize a store containing two keys: "modal1" and "modal2", each mapping to the value "false".
	// When AlpineStoreInit is called, it generates the javascript that binds
	store := make(AlpineStore)

	store[MODAL1] = "false"
	store[MODAL2] = "false"

	func() Node {
		return AppLayout("User Interface Elements", *identity,
			H1(Class("text-4xl font-bold"), Text("Modals")),

			store.Init(),

			ModalActuator(MODAL1, ButtonBlue(Text("Open Modal 1"))),
			ModalActuator(MODAL2, ButtonRed(Text("Open Modal 2"))),

			Modal(
				MODAL1,
				"Test Modal",
				Text("Hello!"),
				[]Node{
					ButtonGray(Text("Close")),
				},
			),

			Modal(
				MODAL2,
				"Another Modal",
				Text("This is the second modal!"),
				[]Node{
					ButtonGray(Text("Close")),
					ButtonRed(Text("Click this!")),
					ButtonBlue(Text("You can as many bottom buttons as you want!")),
				},
			),

			ButtonGray(
				Text("Swap contents of Modal 1 via HTMX"),
				Attr("hx-get", pageinfo.Root.App.Examples.Lipsum_hx.Url()),
				Attr("hx-target", CSSID(MODAL1)),
			),
		)
	}().Render(w)
}
