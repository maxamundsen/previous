package examples

import (
	. "maragu.dev/gomponents"
	. "previous/components"
	. "previous/basic"
	
	"net/http"
	"strconv"
)

func HtmxCounterController(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	count += 1

	CounterButton(count).Render(w)
}

func CounterButton(count int) Node {
	if count == 10 {
		return ButtonBlue(Text("Counter reached 10"))
	} else {
		return ButtonGray(
			Attr("hx-get", "/app/examples/htmx-counter?count="+ToString(count)),
			Attr("hx-swap", "outerHTML"),
			Text("Counter: "), ToText(count),
		)
	}
}
