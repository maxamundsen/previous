package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	. "webdawgengine/basic"
	"webdawgengine/middleware"
	"webdawgengine/models"

	"net/http"
	"strconv"
)

func HtmxController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	HtmxView(*identity).Render(w)
}

func HtmxView(identity models.Identity) Node {
	return AppLayout("HTMX Example", identity,
		P(Class("mb-5"), Text("Click the button to increase the counter. Open up the network tab in the browser developer tools to see how this works under the hood.")),
		CounterButton(0),
		P(Class("mt-5 mb-5"), Text("HTMX is all about sending partial HTML snippets over the network, and swapping that response into the DOM. This button is dynamically generated on the server, and HTMX automatically patches the DOM with the response. Unlike other AJAX methods, you declare the behavior entirely using HTML attributes, which follows LoB principles, and makes using HTMX very ergonomic.")),
		P(Class("mb-5"),
			Text("To learn more about HTMX, click "), PageLink("https://htmx.org", Text("here"), true), Text("."),
			Text(" This example only demonstrates functionality, and not utility. Check out the examples on the HTMX website for better uses of HTMX."),
		),
	)
}

func HtmxCounterController(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.PathValue("count"))
	count += 1

	CounterButton(count).Render(w)
}

func CounterButton(count int) Node {
	if count == 10 {
		return FormInput(Type("text"), Value("I just turned into a form input! What?!"))
	} else {
		return ButtonRed(
			Attr("hx-get", "/app/examples/htmx/counter/"+ToString(count)),
			Attr("hx-swap", "outerHTML"),
			Text("Counter: "), ToText(count),
		)
	}
}
