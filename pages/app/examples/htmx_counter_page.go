package examples

import (
	"previous/.metagen/pageinfo"
	. "previous/basic"
	. "previous/components"

	. "maragu.dev/gomponents"

	"net/http"
	"strconv"
)

func HtmxCounterPage(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	count += 1

	CounterButton(count).Render(w)
}

func CounterButton(count int) Node {
	if count == 10 {
		return ButtonBlue(Icon(ICON_HTMX, 24), Text("Counter reached 10"))
	} else {
		return ButtonGray(
			Attr("hx-get", pageinfo.APP_EXAMPLES_HTMXCOUNTER.Url() + "?count="+ToString(count)),
			Attr("hx-swap", "outerHTML"),
			Text("Counter: "), ToText(count),
		)
	}
}
