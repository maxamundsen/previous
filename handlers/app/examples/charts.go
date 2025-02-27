package examples

import (
	. "previous/components"
	. "previous/handlers/app"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/middleware"

	"net/http"
)

func ChartHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	AppLayout("Chart.js Example", *identity,
		BarChart(),
	).Render(w)
}

func BarChart() Node {
	label_string := MakeJsArray(
		[]string{"test", "hello", "item3", "spiral"},
	)

	data_string := MakeJsArray(
		[]int{3, 5, 8, 13},
	)

	return Card(
		Canvas(),
		InlineScriptf(`
			let ctx = me("canvas", me());

			new Chart(ctx, {
				type: 'bar',
				data: {
					labels: %s,
					datasets: [{
						data: %s,
						borderWidth: 1
					}]
				},
			});`,
			label_string,
			data_string,
		),
	)
}
