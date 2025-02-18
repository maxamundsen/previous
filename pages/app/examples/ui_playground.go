package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	hx "maragu.dev/gomponents-htmx"
	"net/http"
	. "previous/components"
	"previous/middleware"
	. "previous/pages/app"
)

func UIPlaygroundPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	const (
		MODAL1 = "modal1"
		MODAL2 = "modal2"
	)

	func() Node {
		return AppLayout("User Interface Elements", *identity,
			H1(Class("text-4xl font-bold"), Text("Modals")),

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
				hx.Get("/app/examples/lipsum-hx"),
				hx.Target(CSSID(MODAL1)),
			),
		)
	}().Render(w)
}
