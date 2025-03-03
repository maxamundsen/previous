package examples

import (
	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
	"net/http"
	. "previous/ui"
	"previous/middleware"
)

func UIPlaygroundHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	const (
		MODAL1 = "modal1"
		MODAL2 = "modal2"
	)

 	AppLayout("User Interface Elements", LAYOUT_SECTION_EXAMPLES, *identity, session,
		H1(Class("text-4xl font-bold"), Text("Modals")),

		Modal(
			MODAL1,
			Text("This is a modal!"),
			Text(LOREM_IPSUM),
			[]Node{
				ButtonUISuccess(Text("OK")),
				ButtonUI(Text("Close")),
			},
		),

		Modal(
			MODAL2,
			Text("Modal"),
			Text("This is the inside"),
			nil,
		),

		Div(InlineStyle("$me { display: flex; flex-direction: row; gap: $2;}"),
			ModalActuator(MODAL1, ButtonUI(Text("Open Modal 1"))),
			ModalActuator(MODAL1, ButtonUI(Text("Open Modal 1 (secondary actuator)"))),
			ModalActuator(MODAL2, ButtonUI(Text("Open Modal 2"))),
		),

		Br(),

		ButtonUI(
			Text("Swap contents of Modal 1 via HTMX"),
			hx.Get("/app/examples/autotable-hx"),
			hx.Target(CSSID(MODAL1)),
		),

		Divider(),
	).Render(w)
}
