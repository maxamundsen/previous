package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"net/http"
	"strconv"
)

func HtmxController(w http.ResponseWriter, r *http.Request) {
	HtmxView().Render(w)
}

func HtmxView() Node {
	return AppLayout("HTMX Example",
		P(Text("Click the button to increase the counter")),
		CounterButton(0),
	)
}

func HtmxCounterController(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.PathValue("count"))
	count += 1

	CounterButton(count).Render(w)
}

func CounterButton(count int) Node {
	return Button(
		Class("btn btn-primary"),
		Attr("hx-get", "/app/examples/htmx/counter/"+ToString(count)),
		Attr("hx-swap", "outerHTML"),
		Text("Counter: "), ToText(count),
	)
}
