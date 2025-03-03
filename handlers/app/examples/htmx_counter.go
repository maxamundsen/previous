package examples

import (
	. "previous/basic"
	. "previous/ui"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"

	"net/http"
	"strconv"
)

func HtmxCounterHandler(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.PathValue("count"))
	count += 1

	CounterButton(count).Render(w)
}

func CounterButton(count int) Node {
	if count == 10 {
		return ButtonUI(Icon(ICON_HTMX, 24), Text("Counter reached 10"))
	} else {
		return ButtonUI(
			hx.Get("/app/examples/htmx-counter/"+ToString(count)),
			hx.Swap("outerHTML"),
			Text("Counter: "), ToText(count),
		)
	}
}
