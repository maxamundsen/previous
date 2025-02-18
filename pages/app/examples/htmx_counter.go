package examples

import (
	. "previous/basic"
	. "previous/components"

	. "maragu.dev/gomponents"

	"net/http"
	"strconv"
)

func HtmxCounterPage(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.PathValue("count"))
	count += 1

	CounterButton(count).Render(w)
}

func CounterButton(count int) Node {
	if count == 10 {
		return ButtonBlue(Icon(ICON_HTMX, 24), Text("Counter reached 10"))
	} else {
		return ButtonGray(
			Attr("hx-get", "/app/examples/htmx-counter/"+ToString(count)),
			Attr("hx-swap", "outerHTML"),
			Text("Counter: "), ToText(count),
		)
	}
}
